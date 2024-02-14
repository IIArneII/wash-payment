package entity

type Exchange string

const (
	AdminsExchange    Exchange = "admins_exchange"    // Получение рассылки обновления и создания сущностей
	WashBonusExchange Exchange = "wash_bonus_service" // ТОЛЬКО для отправки запроса о том, что нужно запустить рассылку
	PaymentExchange   Exchange = "payment_exchange"   // Для отправки запроса на списание денег и результата списания денег
)
