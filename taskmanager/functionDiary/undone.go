package functionDiary

import (
	"fmt"

	"log"
	"strconv"
)

func Undone(args ...string) string {
	var taskD string
	tasks, readErr := Reader()
	nowtime := TimeFormat()
	if readErr != nil {
		log.Fatal("could not read file : undone.go: ln 42", readErr)
	}
	index, _ := strconv.Atoi(args[0])
	if index > 0 && index <= len(tasks) {
		tasks[index-1].EditToUndone()
		tasks[index-1].Edited = nowtime
		//taskD = tasks[index-1].Status
		err := Writer(tasks)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf(" Task status changed, Edited: %v  \n", nowtime)
	return taskD
}
