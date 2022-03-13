package functionDiary

import (
	"fmt"
	"log"
)

func List(args ...string) {
	tasks, readerErr := Reader()
	if readerErr != nil {
		log.Fatal("fail to read file: list.go: ln 39", readerErr)
	}
	if len(tasks) > 0 {
		//templat := tabwriter.NewWriter(os.Stdout, 7, 1, 1, ' ', 0)
		var count = 0
		for i := 0; i < len(tasks); i++ {
			tasks[i].Serial = i + 1
			if tasks[i].Status == "undone" {
				count++
				var created = " created: "
				var edited = " edited: "
				//fmt.Fprintf(templat, strconv.Itoa(task.Serial)+" "+command.Capitalize(task.Event)+" "+task.Time+"\n")
				fmt.Printf("%-3v%-30v%-11v%-20v%-11v%-20v\n",
					count, Capitalize(tasks[i].Event), created, tasks[i].Created, edited, tasks[i].Edited)
			}
		}

	} else {
		fmt.Println("No Task to Display ")
	}
}
