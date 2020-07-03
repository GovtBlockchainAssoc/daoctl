package views

import (
	"strconv"

	"github.com/GovtBlockchainAssoc/daoctl/models"
	"github.com/alexeyco/simpletable"
)

func paymentHeader() *simpletable.Header {
	return &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Date"},
			{Align: simpletable.AlignCenter, Text: "Account Name"},
			{Align: simpletable.AlignCenter, Text: "Amount"},
			{Align: simpletable.AlignCenter, Text: "Memo"},
		},
	}
}

// PaymentTable is a simpleTable.Table object with payouts
func PaymentTable(payments []models.Payment) *simpletable.Table {

	table := simpletable.New()
	table.Header = paymentHeader()

	for index := range payments {

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: strconv.Itoa(int(payments[index].ID))},
			{Align: simpletable.AlignRight, Text: payments[index].PaymentDate.Time.Format("2006 Jan 02")},
			{Align: simpletable.AlignRight, Text: string(payments[index].Recipient)},
			{Align: simpletable.AlignRight, Text: payments[index].Amount.String()},
			{Align: simpletable.AlignRight, Text: payments[index].Memo},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	return table
}
