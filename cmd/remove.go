package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		storage, err := GetStorage()
		if err != nil {
			panic(err)
		}

		quotes := storage.GetAllQuotes()

		for _, q := range quotes {
			if strings.Contains(q.Text, args[0]) {
				success := storage.RemoveQuote(q)

				if success {
					_, _ = os.Stdout.WriteString("'" + q.Text + "' Removed!")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
