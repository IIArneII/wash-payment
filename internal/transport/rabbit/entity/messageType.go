package entity

type MessageType string

const (
	UserMessageType         MessageType = "admin_service/admin_user"
	OrganizationMessageType MessageType = "admin_service/organization"
	GroupMessageType        MessageType = "admin_service/server_group"
	DataMessageType         MessageType = "admin_service/all_data"
	TransactionMessageType  MessageType = "admin_service/transaction"
)
