package cmd

import (
	"context"
	"fmt"

	"github.com/GovtBlockchainAssoc/daoctl/models"
	"github.com/GovtBlockchainAssoc/daoctl/views"
	"github.com/alexeyco/simpletable"
	"github.com/eoscanada/eos-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getProposalsCmd = &cobra.Command{
	Use:   "proposals [account name]",
	Short: "retrieve proposals",
	Long:  "retrieve all proposals",
	// Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		api := eos.New(viper.GetString("EosioEndpoint"))
		ctx := context.Background()

		var tableScope string
		if viper.GetString("get-proposals-cmd-scope") == "open" {
			tableScope = "proposal"
		} else {
			tableScope = viper.GetString("get-proposals-cmd-scope")
		}

		printProposalTable(ctx, api, "Proposals", tableScope)
	},
}

func printProposalTable(ctx context.Context, api *eos.API, title, scope string) {
	fmt.Println("\n", title, " --scope=", scope)
	proposals := models.Proposals(ctx, api, scope)
	proposalsTable := views.ProposalTable(proposals)
	proposalsTable.SetStyle(simpletable.StyleCompactLite)
	fmt.Println("\n" + proposalsTable.String() + "\n\n")
}

func init() {
	getProposalsCmd.Flags().StringP("scope", "s", "open", "table scope for listing proposals (try 'passedprops')")

	getCmd.AddCommand(getProposalsCmd)
}
