/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goquote",
	Short: "Brighten up your terminal with a cringe-esque quote!",
	Long: `Brighten up your terminal with a cringe-esque sentence.
One Quote at a time!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: showQuote,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goquote.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showQuote(cmd *cobra.Command, args []string) {
	loadEnv()
	storage, err := NewStorage(os.Getenv("STORAGE_FILE_PATH") + os.Getenv("STORAGE_FILE_NAME"))
	if err != nil {
		panic(err)
	}
	quotes, err := storage.GetAllQuotes()
	if err != nil {
		panic(err)
	}
	println(quotes)
}

func loadEnv() {
	fmt.Println("loading env...")
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") //probably should use sops
	}
}
