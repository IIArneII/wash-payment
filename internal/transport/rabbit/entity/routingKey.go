package entity

type RoutingKey string

const (
	WashPaymentRoutingKey RoutingKey = "wash_payment"
	WashBonusRoutingKey   RoutingKey = "wash_bonus"
)
