package cmd

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/hypha-dao/daoctl/util"

	eos "github.com/eoscanada/eos-go"
	"github.com/hypha-dao/daoctl/models"
	"github.com/hypha-dao/daoctl/views"
	"github.com/leekchan/accounting"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getBallotCmd = &cobra.Command{
	Use:   "ballot [ballot name]",
	Short: "retrieve ballot details",
	Long:  "retrieve the ballot times, voters, voting selections, and quorum info",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		api := eos.New(viper.GetString("EosioEndpoint"))
		ctx := context.Background()
		ac := accounting.NewAccounting("", 0, ",", ".", "%s %v", "%s (%v)", "%s --") // TODO: make this configurable

		ballotName := eos.Name(viper.GetString("BallotPrefix" + args[0])) // TODO: this will break; need to make it dynamic

		ballot, err := models.NewBallot(ctx, api, ballotName)
		if err != nil {
			panic("Cannot read ballot: " + args[0])
		}

		fmt.Println("\n\n" + views.BallotHeader(*ballot) + "\n\n")
		votesTable, totalVotes := views.VotesTable(ballot.Votes)
		fmt.Println(votesTable.String())
		voteTokenSupply, err := models.GetVoteTokenSupply(ctx, api)
		if err != nil {
			fmt.Println("Cannot read Vote Token supply.")
			return
		}

		supply := big.NewFloat(float64(voteTokenSupply.Amount) / math.Pow10(int(voteTokenSupply.Precision)))
		votes := big.NewFloat(float64(totalVotes.Amount) / math.Pow10(int(voteTokenSupply.Precision)))
		quorum := supply.Mul(supply, big.NewFloat(0.2))

		var quorumMet, isPassing, isVotingClosed bool
		quorumMet = false
		isPassing = false
		isVotingClosed = false
		quorumFlag := votes.Cmp(quorum)
		if quorumFlag > 0 {
			quorumMet = true
		}

		requiredVotes := util.AssetMult(ballot.RejectVotes, big.NewFloat(4))
		if quorumMet && ballot.PassVotes.Amount > requiredVotes.Amount {
			isPassing = true
		}

		if ballot.EndTime.Before(time.Now()) {
			isVotingClosed = true
		}

		fmt.Println()
		output := []string{
			fmt.Sprintf("Vote Token Supply|%v", util.FormatAsset(voteTokenSupply, 0)),
			fmt.Sprintf("Quorum|%v", ac.FormatMoneyBigFloat(quorum)),
			fmt.Sprintf("Votes|%v", ac.FormatMoneyBigFloat(votes)),
			fmt.Sprintln(),
			fmt.Sprintf("Quorum Met?|%v", quorumMet),
			fmt.Sprintf("Vote Passing?|%v", isPassing),
			fmt.Sprintf("Voting Closed?|%v", isVotingClosed),
		}
		fmt.Println(columnize.SimpleFormat(output))
		fmt.Println()
	},
}

func init() {
	getCmd.AddCommand(getBallotCmd)
}
