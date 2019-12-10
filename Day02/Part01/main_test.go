package main

import (
	"reflect"
	"testing"
)

func TestGetInitialState(t *testing.T) {
	stateCase1 := GetInitialState([]int{1, 0, 0, 0, 99})
	expectedStateCase1 := []int{2, 0, 0, 0, 99}

	if !reflect.DeepEqual(stateCase1, expectedStateCase1) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase1, stateCase1)
	}

	stateCase2 := GetInitialState([]int{2, 3, 0, 3, 99})
	expectedStateCase2 := []int{2, 3, 0, 6, 99}

	if !reflect.DeepEqual(stateCase2, expectedStateCase2) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase2, stateCase2)
	}

	stateCase3 := GetInitialState([]int{2, 4, 4, 5, 99, 0})
	expectedStateCase3 := []int{2, 4, 4, 5, 99, 9801}

	if !reflect.DeepEqual(stateCase3, expectedStateCase3) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase3, stateCase3)
	}

	stateCase4 := GetInitialState([]int{1, 1, 1, 4, 99, 5, 6, 0, 99})
	expectedStateCase4 := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

	if !reflect.DeepEqual(stateCase4, expectedStateCase4) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase4, stateCase4)
	}
}
