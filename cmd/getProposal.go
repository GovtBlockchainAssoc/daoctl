package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/GovtBlockchainAssoc/daoctl/models"
	eos "github.com/eoscanada/eos-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getProposalCmd = &cobra.Command{
	Use:   "proposal [proposal id]",
	Short: "retrieve proposal details",
	Long:  "retrieve the detail about a proposal",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		api := eos.New(viper.GetString("EosioEndpoint"))
		ctx := context.Background()

		proposalID, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("Parse error: Proposal id must be a positive integer (uint64)")
			return
		}
		proposal := models.NewProposalByID(ctx, api, proposalID)

		fmt.Println("\n\nDetail for Proposal: ", proposal.Title, "\n")
		fmt.Println(proposal.String())
		fmt.Println()
	},
}

func init() {
	getCmd.AddCommand(getProposalCmd)
}
