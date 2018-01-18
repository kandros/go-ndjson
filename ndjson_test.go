package ndjson

import (
	"os"
	"testing"
)

func TestToJSON(t *testing.T) {
	f, err := os.Open("testdata.log")

	if err != nil {
		panic(err)
	}

	s := ToJSON(f)

	want := `[{"data":"test"}]`

	if s != want {
		t.Fatalf("\ngot:\n%s\nwant:\n %s", s, want)
	}
}
