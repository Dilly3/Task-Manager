package models

type Item struct {
	Serial   int    `json:"serial,omitempty"`
	Event    string `json:"event,omitempty"`
	Status   string `json:"status,omitempty"`
	Created  string `json:"created"`
	Edited   string `json:"edited"`
	DoneTime string `json:"doneTime"`
}
