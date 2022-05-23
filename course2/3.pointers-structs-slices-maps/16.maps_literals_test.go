package main

import (
	"testing"
)

func TestMapsLiterals(t *testing.T) {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Example": Vertex{
			Lat:  30.8371,
			Long: -43.1725,
		},
	}

	// or
	// var m = map[string]Vertex{
	// 	"Example": {
	// 		Lat:  30.8371,
	// 		Long: -43.1725,
	// 	},
	// }

	actual := m["Example"].Lat
	want := 30.8371

	if actual != want {
		t.Errorf("actual %v expected %v", actual, want)
	}
}
