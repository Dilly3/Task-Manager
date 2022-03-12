/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	diary "taskmanager/taskmanager/functionDiary"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "This command changes the status of tasks from undone to done",
	Long: `This Command,Followed by a Number eg; [ undone 5 ]; Undone' the task with serial number 5

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Undone,
}

func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func Undone(cmd *cobra.Command, args []string) {
	tasks, readErr := diary.Reader()
	nowtime := diary.TimeFormat()

	if readErr != nil {
		log.Fatal("could not read file : undone.go: ln 42", readErr)
	}
	index, _ := strconv.Atoi(args[0])
	if index > 0 && index <= len(tasks) {
		tasks[index-1].EditToUndone()
		tasks[index-1].Edited = nowtime
		err := diary.Writer(tasks)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf(" Task status changed, Edited: %v  \n", nowtime)
}
