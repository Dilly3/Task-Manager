/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// clearlistCmd represents the clearlist command
var clearlistCmd = &cobra.Command{
	Use:   "clearlist",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: ClearList,
}

func init() {
	rootCmd.AddCommand(clearlistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ClearList(cmd *cobra.Command, args []string) {

	_, osErr := os.OpenFile("tasks.csv", os.O_TRUNC, 0644)
	if osErr != nil {
		log.Fatal(osErr)
	}
	fmt.Println(" All Tasks Removed")
}
