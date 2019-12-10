package main

import (
	"reflect"
	"testing"
)

func TestParsingInputs(t *testing.T) {
	actual := parseInputToPaths([]string{"R990,D362", "L995,D598,R577"})
	expected := [][]wire{[]wire{wire{direction: 'R', length: 990}, wire{direction: 'D', length: 362}}, []wire{wire{direction: 'L', length: 995}, wire{direction: 'D', length: 598}, wire{direction: 'R', length: 577}}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

// func TestGetSteps(t *testing.T) {
// 	actual := getSteps([]wire{wire{direction: 'R', length: 3}, wire{direction: 'D', length: 2}})
// 	expected := []point{point{1, 0}, point{2, 0}, point{3, 0}, point{3, 1}, point{3, 2}}
// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Errorf("Expected: %v, Actual: %v", expected, actual)
// 	}
// }
