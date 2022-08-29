package functiondiary

import (
	"fmt"
	"github.com/dilly3/task-manager/models"
	"github.com/dilly3/task-manager/utils"
	"log"
	"os"
)

func CleanUp(_ ...string) {
	var unDoneTasks []models.Item
	tasks, readErr := utils.Reader()
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
	_, osErr := os.OpenFile("tasks.json", os.O_TRUNC, 0644)
	if osErr != nil {
		log.Fatal(osErr)
	}

	writeErr := utils.Writer(unDoneTasks)
	if writeErr != nil {
		log.Fatal(writeErr)
	}

	fmt.Println(" All done Tasks Removed")
}
