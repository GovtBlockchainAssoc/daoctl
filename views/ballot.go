package views

import (
	"fmt"

	"github.com/GovtBlockchainAssoc/daoctl/util"
	"github.com/spf13/viper"

	"github.com/GovtBlockchainAssoc/daoctl/models"
	"github.com/alexeyco/simpletable"
	"github.com/eoscanada/eos-go"
	"github.com/ryanuber/columnize"
)

// BallotHeader ...
func BallotHeader(ballot models.Ballot) string {
	output := []string{
		fmt.Sprintf("Ballot ID||%v", string(ballot.BallotName)),
		fmt.Sprintf("Title||%v", string(ballot.Title)),
		fmt.Sprintf("Status||%v", string(ballot.Status)),
		fmt.Sprintf("Begin Time||%v", ballot.BeginTime.Time.Format("2006 Jan 02 15:04:05")),
		fmt.Sprintf("End Time||%v", ballot.EndTime.Time.Format("2006 Jan 02 15:04:05")),
		fmt.Sprintf("Description||%v", ballot.Description),
	}
	return columnize.SimpleFormat(output)
}

func votesHeader() *simpletable.Header {
	return &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Voter"},
			{Align: simpletable.AlignCenter, Text: "Votes For"},
			{Align: simpletable.AlignCenter, Text: "Votes Against"},
			{Align: simpletable.AlignCenter, Text: "Voting Time"},
		},
	}
}

// VotesTable ...
func VotesTable(votes []models.Vote) (*simpletable.Table, eos.Asset) {

	table := simpletable.New()
	table.Header = votesHeader()

	voteTokenSymbol := viper.GetString("VoteTokenSymbol")
	votesAgainstTotal, _ := eos.NewAssetFromString("0.00 " + voteTokenSymbol)
	votesForTotal, _ := eos.NewAssetFromString("0.00 " + voteTokenSymbol)
	totalVotes, _ := eos.NewAssetFromString("0.00 " + voteTokenSymbol)

	for _, vote := range votes {

		var votesFor, votesAgainst eos.Asset
		if vote.WeightedVotes[0].Key == "pass" {
			votesFor = vote.VotingPower
			votesForTotal = votesForTotal.Add(vote.VotingPower)
		} else {
			votesAgainst = vote.VotingPower
			votesAgainstTotal = votesAgainstTotal.Add(vote.VotingPower)
		}
		totalVotes = totalVotes.Add(vote.VotingPower)

		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: string(vote.Voter)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&votesFor, 0)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&votesAgainst, 0)},
			{Align: simpletable.AlignRight, Text: vote.VotingTime.Time.Format("2006 Jan 02 15:04:05")},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: "Total"},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&votesForTotal, 0)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&votesAgainstTotal, 0)},
			{},
		},
	}

	return table, totalVotes
}
