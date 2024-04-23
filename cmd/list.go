package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var Search string

const SearchFlag = "search"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		storage, err := GetStorage()
		if err != nil {
			panic(err)
		}

		flag, _ := cmd.Flags().GetString(SearchFlag)
		var quotes []Quote

		if flag == "" {
			quotes = storage.GetAllQuotes()
			if len(quotes) == 0 {
				os.Stdout.WriteString("No quotes found")
			}
		} else {
			quotes, err = storage.GetQuoteMatching(flag)
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
