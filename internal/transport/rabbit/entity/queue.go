package entity

type Queue string

const (
	DataQueue              Queue = "data_queue"               // Получение обновленных данных
	WithdrawalRequestQueue Queue = "withdrawal_request_queue" // Получение запроса на оплату
)
