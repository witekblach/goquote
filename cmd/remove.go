package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
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

		quotes, err := storage.GetAllQuotes()
		if err != nil {
			slog.Info("I KNOW WHATS WRONG WITH IT. IT AIN'T GOT NO GAS IN IT")
		}

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
}
