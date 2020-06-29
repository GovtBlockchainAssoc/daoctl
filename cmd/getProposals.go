package cmd

import (
	"context"
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/eoscanada/eos-go"
	"github.com/hypha-dao/daoctl/models"
	"github.com/hypha-dao/daoctl/views"
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

		tableScope := viper.GetString("DAOContract")
		if viper.GetString("get-proposal-cmd-scope") != "open" {
			tableScope = viper.GetString("get-proposal-cmd-scope")
		}

		printProposalTable(ctx, api, "Proposals", tableScope)
	},
}

func printProposalTable(ctx context.Context, api *eos.API, title, scope string) {
	fmt.Println("\n", title)
	proposals := models.Proposals(ctx, api, scope)
	proposalsTable := views.ProposalTable(proposals)
	proposalsTable.SetStyle(simpletable.StyleCompactLite)
	fmt.Println("\n" + proposalsTable.String() + "\n\n")
}

func init() {
	getProposalsCmd.Flags().StringP("scope", "s", "open", "table scope for listing proposals (try 'passedprops')")

	getCmd.AddCommand(getProposalsCmd)
}
