package entity

type Queue string

const (
	PaymentDataQueue       Queue = "payment_admin_data_queue"        // Получение обновленных данных
	PaymentUpdateDataQueue Queue = "payment_admin_data_update_queue" //Только для получения БД
)

// Косяк в названиях из-зи Рудольфа!
