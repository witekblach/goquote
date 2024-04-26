package cmd

import (
	"fmt"
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
			quotes, _ = storage.GetAllQuotes()
			if len(quotes) == 0 {
				os.Stdout.WriteString("No quotes found")
			}
		} else {
			quotes, err = storage.GetQuoteMatching(flag)
			if err != nil {
				panic(err)
			}
		}

		for i, q := range quotes {
			os.Stdout.WriteString(fmt.Sprintf("%d) Viewed: %d times. Quote: %s\n", i+1, q.Viewed, q.Text))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&Search, SearchFlag, "s", "", "Search for a substring")
}
