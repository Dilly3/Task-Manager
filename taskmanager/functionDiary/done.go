package functionDiary

import (
	"fmt"
	"log"
	"strconv"
)

func Done(args ...string) string {
	var event1 string
	var taskD string
	tasks, readErr := Reader()

	var index int
	var nowTime = TimeFormat()
	if readErr != nil {
		log.Fatal("Didnt read file: done.go, Ln 43", readErr)
	}

	index, _ = strconv.Atoi(args[0])
	if (index > 0) && (index <= len(tasks)) {
		tasks[index-1].EditToDone()
		for _, task := range tasks {
			if task.Serial == index {

				task.EditToDone()
				event1 = task.Event
				taskD = task.Status

			}
		}
		writerErr := Writer(tasks)
		if writerErr != nil {
			log.Fatal(writerErr)
		}

		fmt.Printf("%v Done %s\n", event1, nowTime)
	}
	return taskD

}
