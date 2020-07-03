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

var getMembersCmd = &cobra.Command{
	Use:   "members",
	Short: "retrieve members",
	Long:  "retrieve all members and their balances",
	// Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		api := eos.New(viper.GetString("EosioEndpoint"))
		ctx := context.Background()

		printMemberTable(ctx, api, "DAO Members")
	},
}

func printMemberTable(ctx context.Context, api *eos.API, title string) {
	fmt.Println("\n", title)
	members := models.Members(ctx, api)
	membersTable := views.MemberTable(members)
	membersTable.SetStyle(simpletable.StyleCompactLite)
	fmt.Println("\n" + membersTable.String() + "\n\n")
}

func init() {
	getCmd.AddCommand(getMembersCmd)
}
