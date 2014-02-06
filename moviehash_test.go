package moviehash

import (
	"testing"
)

var testdata = map[string]string{
	"breakdance.avi": "8e245d9679d31e12",
	"dummy.bin":      "61f7751fc2a72bfb",
}

var filename = []string{"breakdance.avi", "dummy.bin"}
var hash = []string{"8e245d9679d31e12", "61f7751fc2a72bfc"}

func TestMoviehash(t *testing.T) {
	for filename, hash := range testdata {
		chash, err := ComputeFileHash(filename)
		if err != nil {
			t.Fatal(err)
		}

		if chash != hash {
			t.Errorf("hash for %s should be %s, was %s", filename, hash, chash)
		}
	}
}
