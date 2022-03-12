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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "This Commands Marks a task as done",
	Long: ` This Command followed by a number ; eg; [ done 5 ], marks the task number 5 as done:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Done,
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Done(cmd *cobra.Command, args []string) {
	tasks, readErr := diary.Reader()
	var index int
	var nowTime = diary.TimeFormat()
	if readErr != nil {
		log.Fatal("Didnt read file: done.go, Ln 43", readErr)
	}
	index, _ = strconv.Atoi(args[0])
	if (index > 0) && (index <= len(tasks)) {
		tasks[index-1].EditToDone()
		writerErr := diary.Writer(tasks)
		if writerErr != nil {
			log.Fatal(writerErr)
		}

		fmt.Printf("%v Done %s\n", tasks[index-1].Event, nowTime)
	}
}
