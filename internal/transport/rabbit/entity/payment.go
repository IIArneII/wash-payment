package entity

type Withdrawal struct {
	OrganizationId string `json:"organizationId"`
	Amount         int64  `json:"amount"`
	Service        string `json:"service"`
}

type WithdrawalResult struct {
	OrganizationId string `json:"organizationId"`
	Amount         int64  `json:"amount"`
	Service        string `json:"service"`
	Status         string `json:"status"`
}
