package functiondiary

import "testing"

func TestDone(t *testing.T) {

	s := "1"
	items := Done(s)

	if items != "done" {
		t.Errorf("expected %v but got %v ", "done", items)
	}

}
