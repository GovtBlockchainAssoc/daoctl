package models

import (
	"context"

	eos "github.com/eoscanada/eos-go"
	"github.com/spf13/viper"
)

// Payment ...
type Payment struct {
	ID          uint64             `json:"payment_id"`
	PaymentDate eos.BlockTimestamp `json:"payment_date"`
	Recipient   eos.Name           `json:"recipient"`
	Amount      eos.Asset          `json:"amount"`
	Memo        string             `json:"memo"`
}

// Payments converts a generic DAO Object to a typed Payout
func Payments(ctx context.Context, api *eos.API) []Payment {
	var payments []Payment
	var request eos.GetTableRowsRequest
	request.Code = viper.GetString("DAOContract")
	request.Scope = viper.GetString("DAOContract")
	request.Table = "payments"
	request.Limit = 1000 // TODO: make dynamic, like a scroll
	request.JSON = true
	response, _ := api.GetTableRows(ctx, request)
	response.JSONToStructs(&payments)

	return payments
}
