package rabbit

import (
	"wash-payment/internal/app"
	"wash-payment/internal/config"
	"wash-payment/internal/transport/rabbit/vo"

	"github.com/go-openapi/runtime"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
)

type RabbitService interface {
	SendMessage(msg interface{}, service vo.Service, routingKey vo.RoutingKey, messageType vo.MessageType) (err error)
}

type rabbitService struct {
	l    *zap.SugaredLogger
	conn *amqp.Connection

	washBonusPub    *rabbitmq.Publisher
	washBonusSvcSub *rabbitmq.Consumer
	rabbitSvc       app.RabbitService

	intApi     *client.RabbitIntAPI
	intApiAuth runtime.ClientAuthInfoWriter
}

func New(l *zap.SugaredLogger, cfg config.RabbitMQConfig, rabbitSvc app.RabbitService) (svc *RabbitService, err error) {
	return nil, nil
}
