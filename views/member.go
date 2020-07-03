package views

import (
	"github.com/GovtBlockchainAssoc/daoctl/util"
	"github.com/spf13/viper"

	"github.com/GovtBlockchainAssoc/daoctl/models"
	"github.com/alexeyco/simpletable"
	"github.com/eoscanada/eos-go"
)

func memberHeader() *simpletable.Header {
	return &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Account Name"},
			{Align: simpletable.AlignCenter, Text: viper.GetString("VoteTokenSymbol")},
			{Align: simpletable.AlignCenter, Text: viper.GetString("RewardTokenSymbol")},
			// {Align: simpletable.AlignCenter, Text: "Created Date"},  TODO: add join date to members?
		},
	}
}

// MemberTable is a simpleTable.Table object with payouts
func MemberTable(members []models.Member) *simpletable.Table {

	table := simpletable.New()
	table.Header = memberHeader()

	voteTokenTotal, _ := eos.NewAssetFromString("0.00 " + viper.GetString("VoteTokenSymbol"))
	rewardTokenTotal, _ := eos.NewAssetFromString("0.00 " + viper.GetString("RewardTokenSymbol"))

	for index := range members {

		voteTokenTotal = voteTokenTotal.Add(members[index].VoteTokenBalance)
		rewardTokenTotal = rewardTokenTotal.Add(members[index].RewardTokenBalance)

		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: string(members[index].Account)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&members[index].VoteTokenBalance, 2)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&members[index].RewardTokenBalance, 2)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: "Subtotal"},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&voteTokenTotal, 2)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&rewardTokenTotal, 2)},
		},
	}

	return table
}
