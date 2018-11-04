package random

import (
	"bytes"
	"testing"
)

func TestMapResults(t *testing.T) {
	results := make([]string, 5)
	buffer := bytes.NewBufferString("a\nlong\nstring\nof\nnewline\nseparated\nwords")
	err := mapResults(
		func(word string) {
			results = append(results, word)
		},
		func(word string) bool {
			return len(word) <= 2
		},
		buffer,
	)
	if err != nil {
		t.Fatal(err)
	}
	var found_a bool
	var found_of bool
	var found_long bool
	for _, word := range results {
		if word == "a" {
			found_a = true
		}
		if word == "of" {
			found_of = true
		}
		if word == "long" {
			found_long = true
		}
	}
	if found_a {
		t.Error("'a' was not filtered")
	}
	if found_of {
		t.Error("'of' was not filtered")
	}
	if !found_long {
		t.Error("'long' was filtered")
	}
}
