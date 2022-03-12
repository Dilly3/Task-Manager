package functionDiary

import (
	"testing"
)

func TestAdd(t *testing.T) {

	s := "boy"
	items := Add(s)

	if items[0].Event != s {
		t.Errorf("expected %v but got %v ", s, items[0].Event)
	}

}
