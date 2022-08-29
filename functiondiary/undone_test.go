package functiondiary

import "testing"

func TestUndone(t *testing.T) {

	s := "1"
	status := Undone(s)

	if status != "undone" {
		t.Errorf("expected %v but got %v ", "undone", status)
	}

}
