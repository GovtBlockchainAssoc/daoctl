// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/GovtBlockchainAssoc/daoctl/models"
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func numBytes(input string) int64 {
	numBytes, err := strconv.ParseInt(input, 10, 64)
	errorCheck(fmt.Sprintf("invalid number of bytes %q", input), err)

	if int64(uint32(numBytes)) != numBytes {
		fmt.Printf("Invalid number of bytes: capped at unsigned 32 bits.  That's probably too much RAM anyway.\n")
		return 0
	}
	return numBytes
}

var systemBuyRAMBytesCmd = &cobra.Command{
	Use:   "buyrambytes [payer] [receiver] [num bytes]",
	Short: "Buy RAM at market price, for a given number of bytes.",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()
		ctx := context.Background()
		payer := toAccount(args[0], "payer")
		var actionsPerTrx, actionCounter int
		actionsPerTrx = 20
		actionCounter = 0
		var actions []*eos.Action
		actions = make([]*eos.Action, actionsPerTrx)

		if viper.GetBool("system-buyrambytes-cmd-all-members") {

			numBytes := numBytes(args[1])
			members := models.Members(ctx, api)
			for index, member := range members {

				// if this 'batch' of actions is full or we are at the end of the member list, submit transaction
				if actionCounter >= actionsPerTrx || index >= len(members) {
					// submit transaction
					pushEOSCActionsAndContextFreeActions(ctx, api, nil, actions)
					// reset actions to submit list
					actions = make([]*eos.Action, actionsPerTrx)
					actionCounter = 0
				} else {
					// add another action to this slice of actions
					actions[actionCounter] = system.NewBuyRAMBytes(payer, toAccount(string(member.Account), "member receiver"), uint32(numBytes))
					actionCounter++
				}
			}
		} else {
			receiver := toAccount(args[1], "receiver")
			numBytes := numBytes(args[2])
			pushEOSCActions(ctx, api, system.NewBuyRAMBytes(payer, receiver, uint32(numBytes)))
		}
	},
}

func init() {
	systemCmd.AddCommand(systemBuyRAMBytesCmd)
	systemBuyRAMBytesCmd.Flags().BoolP("all-members", "", false, "buy ram for all DAO members")
}
