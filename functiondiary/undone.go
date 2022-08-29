package functiondiary

import (
	"fmt"

	"github.com/dilly3/task-manager/utils"
	"log"
	"strconv"
)

func Undone(args ...string) string {
	var taskD string
	tasks, readErr := utils.Reader()
	nowtime := utils.TimeFormat()
	if readErr != nil {
		log.Fatal("could not read file : undone.go: ln 42", readErr)
	}
	index, _ := strconv.Atoi(args[0])
	if index > 0 && index <= len(tasks) {
		utils.EditToUndone(&tasks[index-1]).Edited = nowtime

		err := utils.Writer(tasks)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf(" Task status changed, Edited: %v  \n", nowtime)
	return taskD
}
