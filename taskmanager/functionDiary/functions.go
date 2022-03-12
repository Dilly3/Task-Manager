package functionDiary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type Item struct {
	Serial   int    `json:"serial,omitempty"`
	Event    string `json:"event,omitempty"`
	Status   string `json:"status,omitempty"`
	Created  string `json:"created"`
	Edited   string `json:"edited"`
	DoneTime string `json:"doneTime"`
}

func (t *Item) EditToDone() {
	t.Status = "done"

}
func (t *Item) EditToUndone() {
	t.Status = "undone"
	t.Edited = TimeFormat()

}

func Writer(items []Item) error {
	var byteArr []byte
	var marshalErr error
	byteArr, marshalErr = json.Marshal(items)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	ioutilErr := ioutil.WriteFile("tasks.csv", byteArr, 0644)
	if ioutilErr != nil {
		return ioutilErr
	}
	return nil
}
func Reader() ([]Item, error) {
	var byteArr []byte
	var unMarshalErr error
	var ioutilErr error

	byteArr, ioutilErr = ioutil.ReadFile("tasks.csv")
	if unMarshalErr != nil {
		log.Fatal("reader failed to read in Reader Function", ioutilErr)
		return []Item{}, ioutilErr
	}
	var tasks []Item
	if len(byteArr) > 0 {
		if unMarshalErr = json.Unmarshal(byteArr, &tasks); unMarshalErr != nil {
			return []Item{}, unMarshalErr
		}
		//for ind, _ := range tasks {
		//	tasks[ind].Serial = ind + 1
		//}
	}

	return tasks, nil
}

func TempWriter(items []Item) error {
	var byteArr []byte
	var marshalErr error
	byteArr, marshalErr = json.Marshal(items)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	ioutilErr := ioutil.WriteFile("temp.csv", byteArr, 0644)
	if ioutilErr != nil {
		return ioutilErr
	}
	return nil
}

func TempReader() ([]Item, error) {
	var byteArr []byte
	var unMarshalErr error
	var ioutilErr error

	byteArr, ioutilErr = ioutil.ReadFile("temp.csv")
	if unMarshalErr != nil {
		log.Fatal("reader failed to read in Reader Function", ioutilErr)
		return []Item{}, ioutilErr
	}
	var tasks []Item
	if len(byteArr) > 0 {
		if unMarshalErr = json.Unmarshal(byteArr, &tasks); unMarshalErr != nil {
			return []Item{}, unMarshalErr
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
		myTime = fmt.Sprintf("		%v : %02v:%02v:%02vPM", day, hrs, min, sec)
	} else if hrs < 12 {
		myTime = fmt.Sprintf("		%v : %02v:%02v:%02vAM", day, hrs, min, sec)
	} else {
		hrs = 00
		myTime = fmt.Sprintf("		%v : %02v:%02v:%02vAM", day, hrs, min, sec)
	}
	return strings.TrimLeft(myTime, " ")
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
