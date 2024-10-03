package entity

type Exchange string

const (
	AdminsExchange    Exchange = "admins_exchange"
	WashBonusExchange Exchange = "wash_bonus_service"
	PaymentExchange   Exchange = "payment_exchange"
	ControlExchange   Exchange = "control_service"
)
