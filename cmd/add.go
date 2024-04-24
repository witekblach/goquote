package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a quote",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		storage, err := GetStorage()
		if err != nil {
			panic(err)
		}

		success := storage.AddQuote(args[0])

		if success {
			_, _ = os.Stdout.WriteString(args[0] + " - added!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
