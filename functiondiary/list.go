package functiondiary

import (
	"fmt"
	"github.com/dilly3/task-manager/utils"
	"log"
)

func List(_ ...string) {
	tasks, readerErr := utils.Reader()
	if readerErr != nil {
		log.Fatal("fail to read file: list.go: ln 39", readerErr)
	}
	if len(tasks) > 0 {
		var count = 0
		for i := 0; i < len(tasks); i++ {
			tasks[i].Serial = i + 1
			if tasks[i].Status == "undone" {
				count++
				var created = " created: "
				var edited = " edited: "
				//fmt.Fprintf(templat, strconv.Itoa(task.Serial)+" "+command.Capitalize(task.Event)+" "+task.Time+"\n")
				fmt.Printf("%-3v%-30v%-11v%-20v%-11v%-20v\n",
					count, utils.Capitalize(tasks[i].Event), created, tasks[i].Created, edited, tasks[i].Edited)
			}
		}

	} else {
		fmt.Println("No Task to Display ")
	}
}
