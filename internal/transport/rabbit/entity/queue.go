package entity

type Queue string

const (
	DataQueue              Queue = "data_queue"
	WashControlQueue       Queue = "wash_control"
	WithdrawalRequestQueue Queue = "withdrawal_request_queue"
)
