package rabbit

import (
	"encoding/json"
	"errors"
	"wash-payment/internal/transport/rabbit/vo"

	"github.com/wagslane/go-rabbitmq"
)

func (svc *rabbitService) processMessage(d rabbitmq.Delivery) rabbitmq.Action {
	return rabbitmq.NackDiscard
}

func (svc *rabbitService) SendMessage(msg interface{}, service vo.Service, routingKey vo.RoutingKey, messageType vo.MessageType) error {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return nil
	}

	switch service {
	case vo.WashPaymentService:
		return svc.washPaymentPublisher.Publish(
			jsonMsg,
			[]string{string(routingKey)},
			rabbitmq.WithPublishOptionsType(string(messageType)),
			rabbitmq.WithPublishOptionsExchange(string(service)),
		)
	default:
		return errors.New("Unknown service")
	}
}
