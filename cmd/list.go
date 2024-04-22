package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var Search string

const SearchFlag = "search"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		storage, err := GetStorage()
		if err != nil {
			panic(err)
		}

		substring, _ := cmd.Flags().GetString(SearchFlag)
		var quotes []Quote
		if substring == "" {
			quotes = storage.GetAllQuotes()
			if len(quotes) == 0 {
				os.Stdout.WriteString("No quotes found")
			}
		} else {
			quotes, err = storage.GetQuoteMatching(substring)
			if err != nil {
				panic(err)
			}
		}

		for _, q := range quotes {
			os.Stdout.WriteString(q.Text)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&Search, SearchFlag, "s", "", "Search for a substring")
}
