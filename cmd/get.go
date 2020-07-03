package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [options] [type]",
	Short: "retrieve and display objects from the DAO on-chain smart contract",
}

func init() {
	RootCmd.AddCommand(getCmd)
}
