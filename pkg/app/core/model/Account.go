package model

type Account struct {
	Id      int64   `json:"id,omitempty"`
	Number  int64   `json:"number,omitempty"`
	Balance float64 `json:"balance,omitempty"`
	Owner   string  `json:"owner,omitempty"`
	Status  string  `json:"status,omitempty"`
}
