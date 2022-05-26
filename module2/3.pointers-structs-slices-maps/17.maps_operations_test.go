package main

import (
	"testing"
)

func TestMapsOperations(t *testing.T) {

	var m = make(map[string]int)

	// insert element
	m["key1"] = 111

	actual := m["key1"] // retrieve element
	want := 111

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}

	// delete element
	delete(m, "key1")

	actualValue, ok := m["key1"]
	wantOk := false
	var wantValue int

	if ok != wantOk {
		t.Errorf("actual %v expected %v", actual, want)
	}

	if actualValue != wantValue {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
