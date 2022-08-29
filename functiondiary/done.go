package functiondiary

import (
	"fmt"
	"github.com/dilly3/task-manager/utils"
	"log"
	"strconv"
)

func Done(args ...string) string {
	var event1 string
	var taskD string
	tasks, readErr := utils.Reader()

	var index int
	var nowTime = utils.TimeFormat()
	if readErr != nil {
		log.Fatal("Didnt read file: done.go, Ln 43", readErr)
	}

	index, _ = strconv.Atoi(args[0])
	if (index > 0) && (index <= len(tasks)) {
		utils.EditToDone(&tasks[index-1])
		for _, task := range tasks {
			if task.Serial == index {

				utils.EditToDone(&task)
				event1 = task.Event
				taskD = task.Status

			}
		}
		writerErr := utils.Writer(tasks)
		if writerErr != nil {
			log.Fatal(writerErr)
		}

		fmt.Printf("%v Done %s\n", event1, nowTime)
	}
	return taskD

}
