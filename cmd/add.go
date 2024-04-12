package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a quote",
	Args:  cobra.ExactArgs(1),
	Long: `Add a quote but multiline;
Add a quote but multiline;
Add a quote but multiline;

This app adds a quote and displays it.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		fmt.Println("Echo: " + args[0])

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}