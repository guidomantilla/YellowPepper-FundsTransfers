package dto

import (
	"YellowPepper-FundsTransfers/pkg/core/exception"
)

type Transfer struct {
	OriginAccount      int64               `json:"origin_account,omitempty"`
	DestinationAccount int64               `json:"destination_account,omitempty"`
	Amount             float64             `json:"amount,omitempty"`
	Status string              `json:"status,omitempty"`
	Errors exception.Exception `json:"errors,omitempty"`
}
