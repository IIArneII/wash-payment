package entity

type Payment struct {
	OrganizationId string `json:"organizationId"`
	Amount         int64  `json:"amount"`
}
