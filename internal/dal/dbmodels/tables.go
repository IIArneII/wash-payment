package dbmodels

const (
	UsersTable         string = "users"
	OrganizationsTable string = "organizations"
	GroupTable         string = "groups"
	TransactionTable   string = "transactions"
	ServicePriceTable  string = "service_prices"

	ByIDCondition          string = "id = ?"
	ByOrgIDAndSvcCondition string = "organization_id = ? AND service = ?"
	CountSelect            string = "COUNT(*)"
)
