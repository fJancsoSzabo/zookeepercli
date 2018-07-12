package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createrCmd)
	createrCmd.PersistentFlags().StringVar(&acls, "acls", "31", "optional, csv list [1|,2|,4|,8|,16|,31]")

}

var createrCmd = &cobra.Command{
	Use:  "creater",
	RunE: createrExecute,
}

func createrExecute(cmd *cobra.Command, args []string) error {
	force = true
	return createExecute(cmd, args)
}