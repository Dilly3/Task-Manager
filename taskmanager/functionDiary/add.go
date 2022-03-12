package functionDiary

import (
	"fmt"
	"log"
)

func Add(args ...string) []Item {
	defer Cleanup()
	if len(args) > 8 {
		fmt.Println("task wont be Added")
		panic("String too long")

	}
	itemsArr, readErr := Reader() // read from csv file

	if readErr != nil { // check for error
		log.Fatal("reader failed to read ", readErr)
	}
	if len(itemsArr) < 1 {
		itemsArr = append(itemsArr, Item{1, CreateString(args), "undone",
			TimeFormat(), " ", " "})
	} else {
		itemsArr = append(itemsArr, Item{len(itemsArr) + 1, CreateString(args), "undone",
			TimeFormat(), " ", " "})
	}

	writerErr := Writer(itemsArr)
	if writerErr != nil {
		log.Fatal(writerErr)
	}
	fmt.Printf("Task Added \n")
	return itemsArr
}
