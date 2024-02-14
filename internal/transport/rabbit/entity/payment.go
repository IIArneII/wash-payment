package entity

type Withdrawal struct {
	OrganizationId string `json:"organizationId"`
	Amount         int64  `json:"amount"`
	Service        string `json:"service"`
}

type WithdrawalSuccess struct {
	OrganizationId string `json:"organizationId"`
	Amount         int64  `json:"amount"`
	Service        string `json:"service"`
}

type WithdrawalFailure struct {
	OrganizationId string `json:"organizationId"`
	Amount         int64  `json:"amount"`
	Service        string `json:"service"`
	Error          string `json:"errors"`
}
