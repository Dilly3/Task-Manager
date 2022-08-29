/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// todCmd represents the tod command
var todCmd = &cobra.Command{
	Use:   "menu",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Menu:\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s",
			"1. go run main.go add <task>: Add a task",
			"2. go run main.go list: list all the tasks in my list",
			"3. go run main.go clear: clear all the tasks in my list",
			"4. exit: Exit the program",
			"5. go run main.go cleanup: delete all task that are marked as done",
			"6. go run main.go clearlist: clear the list",
			"7. help : Help",
			"8. go run main.go undone <task number>:  Mark a task as undone",
			"9. go run main.go done <task number>:  Mark a task as done",
			"10. menu: Show Menu ")

	},
}

func init() {
	rootCmd.AddCommand(todCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
