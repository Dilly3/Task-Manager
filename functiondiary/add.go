package functiondiary

import (
	"fmt"
	"github.com/dilly3/task-manager/models"
	"github.com/dilly3/task-manager/utils"
	"log"
)

func Add(args ...string) []models.Item {
	defer utils.Cleanup()
	if len(args) > 8 {
		fmt.Println("task wont be Added")
		panic("String too long")

	}
	itemsArr, readErr := utils.Reader() // read from csv file

	if readErr != nil { // check for error
		log.Fatal("reader failed to read ", readErr)
	}
	if len(itemsArr) < 1 {
		itemsArr = append(itemsArr, models.Item{Serial: 1, Event: utils.CreateString(args), Status: "undone",
			Created: utils.TimeFormat(), Edited: " ", DoneTime: " "})
	} else {
		itemsArr = append(itemsArr, models.Item{Serial: len(itemsArr) + 1, Event: utils.CreateString(args), Status: "undone",
			Created: utils.TimeFormat(), Edited: " ", DoneTime: " "})
	}

	writerErr := utils.Writer(itemsArr)
	if writerErr != nil {
		log.Fatal(writerErr)
	}
	fmt.Printf("Task Added \n")
	return itemsArr
}
