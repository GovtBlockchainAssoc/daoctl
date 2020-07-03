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

var getPaymentsCmd = &cobra.Command{
	Use:   "payments",
	Short: "retrieve payments from the DAO to members",
	Long:  "retrieve all payments, including payments of both reward and voting tokens, from the DAO to members",
	// Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		api := eos.New(viper.GetString("EosioEndpoint"))
		ctx := context.Background()

		printPaymentsTable(ctx, api, "Payments")
	},
}

func printPaymentsTable(ctx context.Context, api *eos.API, title string) {
	fmt.Println("\n", title)
	payments := models.Payments(ctx, api)
	paymentsTable := views.PaymentTable(payments)
	paymentsTable.SetStyle(simpletable.StyleCompactLite)
	fmt.Println("\n" + paymentsTable.String() + "\n\n")
}

func init() {
	getCmd.AddCommand(getPaymentsCmd)
}
