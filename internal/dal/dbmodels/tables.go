package dbmodels

const (
	UsersTable         string = "users"
	OrganizationsTable string = "organizations"
	GroupsTable        string = "groups"
	WashServersTable   string = "wash_servers"
	TransactionsTable  string = "transactions"
	ServicePricesTable string = "service_prices"

	ByIDCondition          string = "id = ?"
	ByOrgIDAndSvcCondition string = "organization_id = ? AND service = ?"
	CountSelect            string = "COUNT(*)"
)
