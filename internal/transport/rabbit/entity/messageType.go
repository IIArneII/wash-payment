package entity

type MessageType string

const (
	UserMessageType         MessageType = "admin_service/admin_user"
	OrganizationMessageType MessageType = "admin_service/organization"
	GroupMessageType        MessageType = "admin_service/server_group"
	DataMessageType         MessageType = "admin_service/data"

	WithdrawalRequestMessageType MessageType = "payment_service/withdraw_request"
	WithdrawalSuccessMessageType MessageType = "payment_service/withdraw_success"
	WithdrawalFailureMessageType MessageType = "payment_service/withdraw_failure"
)
