package dbmodels

const (
	UsersTable         string = "users"
	OrganizationsTable string = "organizations"
	GroupTable         string = "groups"
	TransactionTable   string = "transactions"

	ByIDCondition string = "id = ?"
	CountSelect   string = "COUNT(*)"
)
