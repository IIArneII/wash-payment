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
	SendMessage(msg interface{}, service entity.Exchange, routingKey entity.RoutingKey, messageType entity.MessageType) error
}

type rabbitService struct {
	l *zap.SugaredLogger

	washPaymentPublisher *rabbitmq.Publisher
	shareConsumer        *rabbitmq.Consumer
	adminConsumer        *rabbitmq.Consumer

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

	svc.washPaymentPublisher, err = rabbitmq.NewPublisher(
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

	svc.shareConsumer, err = rabbitmq.NewConsumer(
		conn,
		svc.processMessage,
		string(entity.PaymentUpdateDataQueue),

		rabbitmq.WithConsumerOptionsExchangeDeclare,

		rabbitmq.WithConsumerOptionsExchangeName(string(entity.WashBonusExchange)),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),

		rabbitmq.WithConsumerOptionsRoutingKey(string(entity.WashPaymentRoutingKey)),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	svc.adminConsumer, err = rabbitmq.NewConsumer(
		conn,
		svc.processMessage,
		string(entity.PaymentDataQueue),

		rabbitmq.WithConsumerOptionsExchangeDeclare,

		rabbitmq.WithConsumerOptionsExchangeName(string(entity.AdminsExchange)),
		rabbitmq.WithConsumerOptionsExchangeKind("fanout"),
		rabbitmq.WithConsumerOptionsRoutingKey(string(entity.PaymentDataQueue)),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	err = svc.SendMessage(nil, entity.WashBonusExchange, entity.WashBonusRoutingKey, entity.DataMessageType)
	if err != nil {
		return nil, err
	}

	return svc, nil
}
