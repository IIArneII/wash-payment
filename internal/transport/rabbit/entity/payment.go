package entity

type Withdrawal struct {
	GroupId   string `json:"groupId"`
	PostСount string `json:"postСount"`
	Amount    int64  `json:"amount"`
	Service   string `json:"service"`
}

type WithdrawalSuccess struct {
	GroupId   string `json:"groupId"`
	PostСount string `json:"postСount"`
	Amount    int64  `json:"amount"`
	Service   string `json:"service"`
}

type WithdrawalFailure struct {
	GroupId   string `json:"groupId"`
	PostСount string `json:"postСount"`
	Amount    int64  `json:"amount"`
	Service   string `json:"service"`
	Error     string `json:"errors"`
}
