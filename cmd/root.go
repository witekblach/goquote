package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goquote",
	Short: "Brighten up your terminal with a cringe-esque quote!",
	Long: `Brighten up your terminal with a cringe-esque sentence.
One Quote at a time!`,
	Run: showQuote,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showQuote(cmd *cobra.Command, args []string) {
	storage, err := GetStorage()
	if err != nil {
		panic(err)
	}

	quote, err := storage.GetRandomQuote()
	if err != nil {
		slog.Info("try adding a quote first :) ")
		return
	}

	os.Stdout.WriteString(quote.Text)
}
