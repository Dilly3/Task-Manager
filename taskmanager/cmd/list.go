/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	diary "taskmanager/taskmanager/functionDiary"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks in my list",
	Long: `:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: List,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func List(cmd *cobra.Command, args []string) {
	tasks, readerErr := diary.Reader()
	if readerErr != nil {
		log.Fatal("fail to read file: list.go: ln 39", readerErr)
	}
	if len(tasks) > 0 {
		//templat := tabwriter.NewWriter(os.Stdout, 7, 1, 1, ' ', 0)
		var count = 1
		for _, task := range tasks {

			if task.Status == "undone" {
				var created = " created: "
				var edited = " edited: "
				//fmt.Fprintf(templat, strconv.Itoa(task.Serial)+" "+command.Capitalize(task.Event)+" "+task.Time+"\n")
				fmt.Printf("%-3v%-30v%-11v%-20v%-11v%-20v\n",
					task.Serial, diary.Capitalize(task.Event), created, task.Created, edited, task.Edited)
				count++
			}

		}
	} else {
		fmt.Println("No Task to Display \n")
	}

}
