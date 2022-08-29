package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dilly3/task-manager/models"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func Writer(items []models.Item) error {
	var byteArr []byte
	var marshalErr error
	byteArr, marshalErr = json.Marshal(items)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	ioutilErr := ioutil.WriteFile("./tasks.json", byteArr, 0644)
	if ioutilErr != nil {
		return ioutilErr
	}
	return nil
}
func Reader() ([]models.Item, error) {
	var byteArr []byte
	var unMarshalErr error
	var ioutilErr error

	byteArr, ioutilErr = ioutil.ReadFile("./tasks.json")
	if unMarshalErr != nil {
		log.Fatal("reader failed to read in Reader Function", ioutilErr)
		return []models.Item{}, ioutilErr
	}
	var tasks []models.Item
	if len(byteArr) > 0 {
		if unMarshalErr = json.Unmarshal(byteArr, &tasks); unMarshalErr != nil {
			return []models.Item{}, unMarshalErr
		}
		//for ind, _ := range tasks {
		//	tasks[ind].Serial = ind + 1
		//}
	}

	return tasks, nil
}

func TempWriter(items []models.Item) error {
	var byteArr []byte
	var marshalErr error
	byteArr, marshalErr = json.Marshal(items)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	ioutilErr := ioutil.WriteFile("temp.json", byteArr, 0644)
	if ioutilErr != nil {
		return ioutilErr
	}
	return nil
}

func TempReader() ([]models.Item, error) {
	var byteArr []byte
	var unMarshalErr error
	var ioutilErr error

	byteArr, ioutilErr = ioutil.ReadFile("temp.json")
	if unMarshalErr != nil {
		log.Fatal("reader failed to read in Reader Function", ioutilErr)
		return []models.Item{}, ioutilErr
	}
	var tasks []models.Item
	if len(byteArr) > 0 {
		if unMarshalErr = json.Unmarshal(byteArr, &tasks); unMarshalErr != nil {
			return []models.Item{}, unMarshalErr
		}
		//for ind, _ := range tasks {
		//	tasks[ind].Serial = ind + 1
		//}
	}

	return tasks, nil
}

//func (i *Item) Label() string {
//	return strconv.Itoa(i.Serial) + " "
//}

func Capitalize(s string) string {
	s = strings.ToLower(s)
	var capString []byte
	for ind := 0; ind < len(s); ind++ {
		if ind == 0 {
			capString = append(capString, s[ind]-32)
		} else {
			capString = append(capString, s[ind])
		}
	}
	return string(capString)
}
func Cleanup() {
	errMessage := recover()
	if errMessage != nil {
		log.Fatal(errMessage)
	}
}

func TimeFormat() string {
	currentDate := time.Now()
	day := currentDate.Weekday()
	hrs := currentDate.Hour()
	sec := currentDate.Second()
	min := currentDate.Minute()
	var myTime string
	if hrs > 12 {
		hrs = hrs - 12
		myTime = fmt.Sprintf("\t\t%10v : %02v:%02v:%02vPM", day, hrs, min, sec)
	} else if hrs < 12 {
		myTime = fmt.Sprintf("\t\t%v : %02v:%02v:%02vAM", day, hrs, min, sec)
	} else {
		hrs = 00
		myTime = fmt.Sprintf("\t\t%v : %02v:%02v:%02vAM", day, hrs, min, sec)
	}
	return strings.TrimLeft(myTime, " ")
}
func EditToUndone(item *models.Item) *models.Item {
	item.Status = "undone"
	item.Edited = TimeFormat()
	return item

}
func EditToDone(item *models.Item) *models.Item {
	item.Status = "done"
	return item

}

func CreateString(args []string) string {
	var space string = " "
	var tasks string
	switch len(args) {
	case 1:
		{
			tasks = args[0]
		}
	case 2:
		{
			tasks = args[0] + space + args[1]
		}
	case 3:
		{
			tasks = args[0] + space + args[1] + space + args[2]

		}
	case 4:
		{
			tasks = args[0] + space + args[1] +
				space + args[2] + space + args[3]
		}
	case 5:
		{
			tasks = args[0] + space + args[1] + space + args[2] + space + args[3] + space + args[4]
		}
	case 6:
		{
			tasks = args[0] + space + args[1] +
				space + args[2] + space + args[3] + space + args[4] + space + args[5]
		}
	case 7:
		{
			tasks = args[0] + space + args[1] +
				space + args[2] + space + args[3] + space + args[4] + space + args[5] + space + args[6]
		}
	case 8:
		{
			tasks = args[0] + space + args[1] +
				space + args[2] + space + args[3] + space + args[4] + space + args[5] + space + args[6] + space + args[7]
		}
	default:
		{
			tasks = args[0] + space + args[1] +
				space + args[2] + space + args[3] + space + args[4] + space + args[5] + space + args[6] +
				space + args[7] + space + args[8]
		}
	}
	return tasks
}
