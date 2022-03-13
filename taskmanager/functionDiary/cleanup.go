package functionDiary

import (
	"fmt"
	"log"
	"os"
)

func CleanUp(args ...string) {
	var unDoneTasks []Item
	tasks, readErr := Reader()
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
	_, osErr := os.OpenFile("tasks.csv", os.O_TRUNC, 0644)
	if osErr != nil {
		log.Fatal(osErr)
	}

	writeErr := Writer(unDoneTasks)
	if writeErr != nil {
		log.Fatal(writeErr)
	}

	fmt.Println(" All done Tasks Removed")
}
