package rabbit

import (
	"context"
	"encoding/json"
	"time"
	"wash-payment/internal/transport/rabbit/entity"

	"github.com/wagslane/go-rabbitmq"
)

func (svc *rabbitService) processMessage(d rabbitmq.Delivery) rabbitmq.Action {
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	svc.l.Info("TYPE -> ", d.Type)

	switch entity.MessageType(d.Type) {
	case entity.OrganizationMessageType:
		var msg entity.Organization
		err := json.Unmarshal(d.Body, &msg)

		if err != nil {
			svc.l.Info(err)
			return rabbitmq.NackRequeue
		}

		err = svc.rabbitSvc.UpsertOrganization(cxt, msg)

		if err != nil {
			svc.l.Info(err)
			return rabbitmq.NackRequeue
		}

	case entity.GroupMessageType:
		var msg entity.Group
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			svc.l.Info(err)
			return rabbitmq.NackRequeue
		}

		err = svc.rabbitSvc.UpsertGroup(cxt, msg)
		if err != nil {
			svc.l.Info(err)
			return rabbitmq.NackRequeue
		}

	case entity.UserMessageType:
		var msg entity.User
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			svc.l.Info(err)
			return rabbitmq.NackRequeue
		}

		err = svc.rabbitSvc.UpsertUser(cxt, msg)
		if err != nil {
			svc.l.Info(err)
			return rabbitmq.NackRequeue
		}

	case entity.WithdrawalMessageType:
		var msg entity.Payment
		err := json.Unmarshal(d.Body, &msg)

		if err != nil {
			svc.l.Info(err)
			return rabbitmq.NackDiscard
		}

		err = svc.rabbitSvc.ProcessWithdrawal(cxt, msg)
		if err != nil {
			svc.l.Info(err)
			return rabbitmq.NackDiscard
		}
		_ = svc.SendMessage(nil, entity.WashBonusExchange, entity.WashPaymentRoutingKey, entity.WithdrawalMessageType)

	default:
		return rabbitmq.NackDiscard
	}

	return rabbitmq.Ack
}

func (svc *rabbitService) SendMessage(msg interface{}, service entity.Exchange, routingKey entity.RoutingKey, messageType entity.MessageType) error {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	switch service {
	case entity.WashBonusExchange:
		return svc.washPaymentPublisher.Publish(
			jsonMsg,
			[]string{string(routingKey)},
			rabbitmq.WithPublishOptionsType(string(messageType)),
			rabbitmq.WithPublishOptionsReplyTo(string(entity.PaymentDataQueue)),
			rabbitmq.WithPublishOptionsExchange(string(service)),
		)
	default:
		panic("Unknown service")
	}
}
