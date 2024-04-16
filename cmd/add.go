package cmd

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/tursodatabase/go-libsql"
	"log"
	"os"
	"path/filepath"
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

		//get connection to db, insert,

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env File") //probably should use sops
		}

		dbName := "local.db"
		primaryUrl := "libsql://" + os.Getenv("TURSO_HOST")
		authToken := os.Getenv("TURSO_TOKEN")

		dir, err := os.MkdirTemp("", "libsql-*")
		if err != nil {
			fmt.Println("Error creating temporary directory:", err)
			os.Exit(1)
		}
		defer os.RemoveAll(dir)

		dbPath := filepath.Join(dir, dbName)

		connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
			libsql.WithAuthToken(authToken),
		)
		if err != nil {
			fmt.Println("Error creating connector:", err)
			os.Exit(1)
		}
		defer connector.Close()

		db := sql.OpenDB(connector)
		defer db.Close()

		rows, err := db.Query("SELECT * FROM quotes")
		if err != nil {
			fmt.Println("AAAAAA PANIC")
			panic(err)
		}
		defer rows.Close()
		var albums []Quote

		// Loop through rows, using Scan to assign column data to struct fields.
		for rows.Next() {
			var alb Quote
			if err := rows.Scan(&alb.Text, &alb.Viewed); err != nil {
				println(fmt.Println(albums, err))
			}
			albums = append(albums, alb)
		}
		if err = rows.Err(); err != nil {
			println(fmt.Println(albums, err))
		}

		println(fmt.Println(albums))
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
