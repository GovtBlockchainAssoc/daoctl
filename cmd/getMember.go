package cmd

import (
	"context"
	"fmt"

	"github.com/GovtBlockchainAssoc/daoctl/models"

	eos "github.com/eoscanada/eos-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getMemberCmd = &cobra.Command{
	Use:   "member [account name]",
	Short: "retrieve member details",
	Long:  "retrieve member details, including balances",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		api := eos.New(viper.GetString("EosioEndpoint"))
		ctx := context.Background()

		account := toName(args[0], "member account")
		member := models.NewMember(ctx, api, account)

		fmt.Println(member.String())
		fmt.Println()
	},
}

func init() {
	getCmd.AddCommand(getMemberCmd)
}
