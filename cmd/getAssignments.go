package cmd

import (
  "context"
  "fmt"
  "github.com/alexeyco/simpletable"
  "github.com/eoscanada/eos-go"
  "github.com/hypha-dao/daoctl/models"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var getAssignmentCmd = &cobra.Command{
  Use:   "assignments [account name]",
  Short: "retrieve assignments",
  Long:  "retrieve all active assignments For a json dump, append the argument --json.",
  Args:  cobra.RangeArgs(1, 2),
  Run: func(cmd *cobra.Command, args []string) {
    api := eos.New("https://api.telos.kitchen")
    ctx := context.Background()

    periods := models.LoadPeriods(api)

    roles := models.Roles(ctx, api, periods)

    assignments := models.Assignments(ctx, api, roles, periods)
    assignmentsTable := models.AssignmentTable(assignments)
    assignmentsTable.SetStyle(simpletable.StyleCompactLite)
    fmt.Println("\n\n" + assignmentsTable.String() + "\n\n")

    if viper.GetBool("include-proposals") == true {
      	propAssignments := models.ProposedAssignments(ctx, api, roles, periods)
      	propAssignmentsTable := models.AssignmentTable(propAssignments)
      	propAssignmentsTable.SetStyle(simpletable.StyleCompactLite)
      	fmt.Println("\n\n" + propAssignmentsTable.String() + "\n\n")
      return
    }
  },
}

func init() {
  getCmd.AddCommand(getAssignmentCmd)
}
