package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long: `Add new task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Added new task: %s", args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
