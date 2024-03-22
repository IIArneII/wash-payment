package rabbit

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"wash-payment/internal/app"
	"wash-payment/internal/transport/rabbit/entity"

	"github.com/wagslane/go-rabbitmq"
)

func (svc *rabbitService) processMessage(d rabbitmq.Delivery) rabbitmq.Action {
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	svc.l.Infof("Message: %s, %s", d.Type, string(d.Body))

	switch entity.MessageType(d.Type) {
	case entity.OrganizationMessageType:
		var msg entity.Organization
		err := json.Unmarshal(d.Body, &msg)

		if err != nil {
			svc.l.Error(err)
			return rabbitmq.NackRequeue
		}

		err = svc.rabbitSvc.UpsertOrganization(cxt, msg)

		if err != nil {
			if errors.Is(err, app.ErrOldVersion) {
				return rabbitmq.NackDiscard
			}
			svc.l.Error(err)
			return rabbitmq.NackRequeue
		}

	case entity.GroupMessageType:
		var msg entity.Group
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			svc.l.Error(err)
			return rabbitmq.NackRequeue
		}

		err = svc.rabbitSvc.UpsertGroup(cxt, msg)
		if err != nil {
			if errors.Is(err, app.ErrOldVersion) {
				return rabbitmq.NackDiscard
			}
			svc.l.Error(err)
			return rabbitmq.NackRequeue
		}

	case entity.UserMessageType:
		var msg entity.User
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			svc.l.Error(err)
			return rabbitmq.NackRequeue
		}

		err = svc.rabbitSvc.UpsertUser(cxt, msg)
		if err != nil {
			if errors.Is(err, app.ErrOldVersion) {
				return rabbitmq.NackDiscard
			}
			svc.l.Error(err)
			return rabbitmq.NackRequeue
		}

	case entity.WithdrawalRequestMessageType:
		var msg entity.Withdrawal
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			svc.l.Error(err)
			return rabbitmq.NackDiscard
		}

		err = svc.rabbitSvc.Withdrawal(cxt, msg)
		if err != nil {
			svc.l.Info(err)
			_ = svc.SendMessage(withdrawalFailureMsg(msg, err), entity.PaymentExchange, d.ReplyTo, entity.WithdrawalFailureMessageType)
			return rabbitmq.NackDiscard
		}
		svc.l.Info("Отправка ответа об успехеcls")
		err = svc.SendMessage(withdrawalMsg(msg), entity.PaymentExchange, d.ReplyTo, entity.WithdrawalSuccessMessageType)
		if err != nil {
			svc.l.Info(err)
		}

	default:
		return rabbitmq.NackDiscard
	}

	return rabbitmq.Ack
}

func (svc *rabbitService) SendMessage(msg interface{}, service entity.Exchange, routingKey string, messageType entity.MessageType) error {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	switch service {
	case entity.WashBonusExchange:
		return svc.washBonusPublisher.Publish(
			jsonMsg,
			[]string{string(routingKey)},
			rabbitmq.WithPublishOptionsType(string(messageType)),
			rabbitmq.WithPublishOptionsReplyTo(string(entity.DataQueue)),
			rabbitmq.WithPublishOptionsExchange(string(entity.WashBonusExchange)),
		)
	case entity.PaymentExchange:
		return svc.paymentPublisher.Publish(
			jsonMsg,
			[]string{string(routingKey)},
			rabbitmq.WithPublishOptionsType(string(messageType)),
			rabbitmq.WithPublishOptionsExchange(string(entity.PaymentExchange)),
		)
	default:
		panic("Unknown service")
	}
}

func withdrawalFailureMsg(withdrawal entity.Withdrawal, err error) entity.WithdrawalFailure {
	return entity.WithdrawalFailure{
		Service:       withdrawal.Service,
		StationsСount: withdrawal.StationsСount,
		ForDate:       withdrawal.ForDate,
		Error:         err.Error(),
	}
}

func withdrawalMsg(withdrawal entity.Withdrawal) entity.WithdrawalSuccess {
	return entity.WithdrawalSuccess{
		Service:       withdrawal.Service,
		StationsСount: withdrawal.StationsСount,
		ForDate:       withdrawal.ForDate,
	}
}
