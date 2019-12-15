package main

import (
	"reflect"
	"testing"
)

func TestPerformDiagnosis(t *testing.T) {
	stateCase1 := PerformDiagnosis([]int{1, 0, 0, 0, 99}, 1)
	expectedStateCase1 := []int{2, 0, 0, 0, 99}

	if !reflect.DeepEqual(stateCase1, expectedStateCase1) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase1, stateCase1)
	}

	stateCase2 := PerformDiagnosis([]int{2, 3, 0, 3, 99}, 1)
	expectedStateCase2 := []int{2, 3, 0, 6, 99}

	if !reflect.DeepEqual(stateCase2, expectedStateCase2) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase2, stateCase2)
	}

	stateCase3 := PerformDiagnosis([]int{2, 4, 4, 5, 99, 0}, 1)
	expectedStateCase3 := []int{2, 4, 4, 5, 99, 9801}

	if !reflect.DeepEqual(stateCase3, expectedStateCase3) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase3, stateCase3)
	}

	stateCase4 := PerformDiagnosis([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 1)
	expectedStateCase4 := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

	if !reflect.DeepEqual(stateCase4, expectedStateCase4) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase4, stateCase4)
	}

	stateCase5 := PerformDiagnosis([]int{3, 0, 4, 0, 99}, 1)
	expectedStateCase5 := []int{1, 0, 4, 0, 99}

	if !reflect.DeepEqual(stateCase5, expectedStateCase5) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase5, stateCase5)
	}

	stateCase6 := PerformDiagnosis([]int{1002, 4, 3, 4, 33}, 1)
	expectedStateCase6 := []int{1002, 4, 3, 4, 99}

	if !reflect.DeepEqual(stateCase6, expectedStateCase6) {
		t.Errorf("Expected: %v, Actual: %v", expectedStateCase6, stateCase6)
	}
}
