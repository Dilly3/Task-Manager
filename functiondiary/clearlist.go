package functiondiary

import (
	"fmt"
	"log"
	"os"
)

func ClearList() {

	_, osErr := os.OpenFile("tasks.csv", os.O_TRUNC, 0644)
	if osErr != nil {
		log.Fatal(osErr)
	}
	fmt.Println(" All Tasks Removed")
}
