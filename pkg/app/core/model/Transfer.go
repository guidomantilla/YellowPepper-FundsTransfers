package model

type Transfer struct {
	Id                 int64   `json:"id,omitempty"`
	OriginAccount      int64   `json:"origin_account,omitempty"`
	DestinationAccount int64   `json:"destination_account,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	Date               string  `json:"date,omitempty"`
	Status             string  `json:"status,omitempty"`
}
