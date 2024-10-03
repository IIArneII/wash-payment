package entity

import "time"

type Withdrawal struct {
	StationsCount int       `json:"stationsCount"`
	Service       string    `json:"service"`
	ForDate       time.Time `json:"forDate"`
	WashServerID  string    `json:"washServerID"`
}

type WithdrawalSuccess struct {
	StationsCount int       `json:"stationsCount"`
	Service       string    `json:"service"`
	ForDate       time.Time `json:"forDate"`
	WashServerID  string    `json:"washServerID"`
}

type WithdrawalFailure struct {
	StationsCount int       `json:"stationsCount"`
	Service       string    `json:"service"`
	ForDate       time.Time `json:"forDate"`
	WashServerID  string    `json:"washServerID"`
	Error         string    `json:"error"`
}
