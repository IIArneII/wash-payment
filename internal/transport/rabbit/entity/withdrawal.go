package entity

import "time"

type Withdrawal struct {
	GroupId       string    `json:"groupId"`
	StationsСount int       `json:"stationsСount"`
	Service       string    `json:"service"`
	ForDate       time.Time `json:"forDate"`
}

type WithdrawalSuccess struct {
	GroupId       string `json:"groupId"`
	StationsСount int    `json:"stationsСount"`
	Service       string `json:"service"`
}

type WithdrawalFailure struct {
	GroupId       string `json:"groupId"`
	StationsСount int    `json:"stationsСount"`
	Service       string `json:"service"`
	Error         string `json:"errors"`
}
