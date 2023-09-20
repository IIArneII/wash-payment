package rabbit

import (
	"wash-payment/internal/app"
	"wash-payment/internal/config"
	"wash-payment/internal/transport/rabbit/vo"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type RabbitService interface {
	SendMessage(msg interface{}, service vo.Service, routingKey vo.RoutingKey, messageType vo.MessageType) error
}

type rabbitService struct {
	l *zap.SugaredLogger

	washPaymentPublisher *rabbitmq.Publisher
	washPaymentConsumer  *rabbitmq.Consumer
	rabbitSvc            app.RabbitService
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
		rabbitmq.WithPublisherOptionsExchangeName(string(vo.WashPaymentService)),
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	svc.washPaymentConsumer, err = rabbitmq.NewConsumer(
		conn,
		svc.processMessage,
		string(vo.WashPaymentRoutingKey),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeName(string(vo.WashPaymentService)),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),
		rabbitmq.WithConsumerOptionsRoutingKey(string(vo.WashPaymentRoutingKey)),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return nil, err
	}

	return svc, nil
}
