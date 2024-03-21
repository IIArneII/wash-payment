package rabbit

import (
	"wash-payment/internal/app"
	"wash-payment/internal/config"
	"wash-payment/internal/transport/rabbit/entity"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type RabbitService interface {
	SendMessage(msg interface{}, service entity.Exchange, routingKey string, messageType entity.MessageType) error
}

type rabbitService struct {
	l *zap.SugaredLogger

	washBonusPublisher *rabbitmq.Publisher
	paymentPublisher   *rabbitmq.Publisher
	adminConsumer      *rabbitmq.Consumer
	paymentConsumer    *rabbitmq.Consumer

	rabbitSvc app.RabbitService
}

func NewRabbitService(l *zap.SugaredLogger, cfg config.RabbitMQConfig, rabbitSvc app.RabbitService) (RabbitService, error) {
	svc := &rabbitService{
		l:         l,
		rabbitSvc: rabbitSvc,
	}

	rabbitConf := rabbitmq.Config{
		SASL: []amqp.Authentication{
			&amqp.PlainAuth{
				Username: cfg.User,
				Password: cfg.Password,
			},
		},
		Vhost:      "/",
		ChannelMax: 0,
		FrameSize:  0,
		Heartbeat:  0,
		Properties: nil,
		Locale:     "",
		Dial:       nil,
	}

	conn, err := rabbitmq.NewConn(
		cfg.DSN(),
		rabbitmq.WithConnectionOptionsLogging,
		rabbitmq.WithConnectionOptionsConfig(rabbitConf),
	)
	if err != nil {
		return nil, err
	}

	svc.washBonusPublisher, err = rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeName(string(entity.WashBonusExchange)),
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	svc.paymentPublisher, err = rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeName(string(entity.PaymentExchange)),
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	svc.adminConsumer, err = rabbitmq.NewConsumer(
		conn,
		svc.processMessage,
		string(entity.DataQueue),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeName(string(entity.AdminsExchange)),
		rabbitmq.WithConsumerOptionsRoutingKey(string(entity.WashBonusRoutingKey)),
		rabbitmq.WithConsumerOptionsExchangeKind("fanout"),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	svc.paymentConsumer, err = rabbitmq.NewConsumer(
		conn,
		svc.processMessage,
		string(entity.WithdrawalRequestQueue),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeName(string(entity.PaymentExchange)),
		rabbitmq.WithConsumerOptionsRoutingKey(string(entity.WithdrawalRequestQueue)),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	err = svc.SendMessage(nil, entity.WashBonusExchange, string(entity.WashBonusRoutingKey), entity.DataMessageType)
	if err != nil {
		return nil, err
	}

	return svc, nil
}
