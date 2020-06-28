package views

import (
	"strconv"

	"github.com/hypha-dao/daoctl/util"
	"github.com/spf13/viper"

	"github.com/alexeyco/simpletable"
	"github.com/eoscanada/eos-go"
	"github.com/hypha-dao/daoctl/models"
)

func proposalHeader() *simpletable.Header {
	return &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Proposer"},
			{Align: simpletable.AlignCenter, Text: "Recipient"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: viper.GetString("VoteTokenSymbol")},
			{Align: simpletable.AlignCenter, Text: viper.GetString("RewardTokenSymbol")},
			{Align: simpletable.AlignCenter, Text: "Created Date"},
			{Align: simpletable.AlignCenter, Text: "Ballot"},
		},
	}
}

// ProposalTable is a simpleTable.Table object with payouts
func ProposalTable(proposals []models.Proposal) *simpletable.Table {

	table := simpletable.New()
	table.Header = proposalHeader()

	voteTokenTotal, _ := eos.NewAssetFromString("0.00 " + viper.GetString("VoteTokenSymbol"))
	rewardTokenTotal, _ := eos.NewAssetFromString("0.00 " + viper.GetString("RewardTokenSymbol"))

	for index := range proposals {

		voteTokenTotal = voteTokenTotal.Add(proposals[index].VoteTokenAmount)
		rewardTokenTotal = rewardTokenTotal.Add(proposals[index].RewardTokenAmount)

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: strconv.Itoa(int(proposals[index].ID))},
			{Align: simpletable.AlignRight, Text: string(proposals[index].Proposer)},
			{Align: simpletable.AlignRight, Text: string(proposals[index].Recipient)},
			{Align: simpletable.AlignLeft, Text: proposals[index].Title},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&proposals[index].VoteTokenAmount, 0)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&proposals[index].RewardTokenAmount, 0)},
			{Align: simpletable.AlignRight, Text: proposals[index].CreatedDate.Time.Format("2006 Jan 02")},
			{Align: simpletable.AlignRight, Text: string(proposals[index].BallotName)[11:]},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{},
			{},
			{},
			{Align: simpletable.AlignRight, Text: "Subtotal"},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&voteTokenTotal, 0)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&rewardTokenTotal, 0)},
			{}, {},
		},
	}

	return table
}
