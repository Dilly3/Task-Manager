/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	diary "taskmanager/taskmanager/functionDiary"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: CleanUp,
}

func init() {
	rootCmd.AddCommand(cleanupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func CleanUp(cmd *cobra.Command, args []string) {
	var unDoneTasks []diary.Item
	tasks, readErr := diary.Reader()
	if readErr != nil {
		log.Fatal(readErr)
	}
	for _, task := range tasks {
		if task.Status == "done" {
			continue
		} else {
			unDoneTasks = append(unDoneTasks, task)
		}
	}
	_, osErr := os.OpenFile("task.csv", os.O_TRUNC, 0644)
	if osErr != nil {
		log.Fatal(osErr)
	}

	writeErr := diary.Writer(unDoneTasks)
	if writeErr != nil {
		log.Fatal(writeErr)
	}

	fmt.Println(" All done Tasks Removed")
}
