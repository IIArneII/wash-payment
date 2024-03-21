package entity

import "time"

type Withdrawal struct {
	GroupId       string    `json:"groupId"`
	StationsСount int       `json:"stationsСount"`
	Service       string    `json:"service"`
	ForDate       time.Time `json:"forDate"`
	WashServerID  string    `json:"washServerID"`
}

type WithdrawalSuccess struct {
	GroupId       string    `json:"groupId"`
	StationsСount int       `json:"stationsСount"`
	Service       string    `json:"service"`
	ForDate       time.Time `json:"forDate"`
	WashServerID  string    `json:"washServerID"`
}

type WithdrawalFailure struct {
	GroupId       string    `json:"groupId"`
	StationsСount int       `json:"stationsСount"`
	Service       string    `json:"service"`
	ForDate       time.Time `json:"forDate"`
	WashServerID  string    `json:"washServerID"`
	Error         string    `json:"error"`
}
