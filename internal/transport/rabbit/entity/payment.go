package entity

type Payment struct {
	Organization Organization `json:"organization"`
	Amount       int64        `json:"amount"`
}
