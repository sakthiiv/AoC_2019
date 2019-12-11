package main

import (
	"reflect"
	"testing"
)

func TestEvaluateCriteria(t *testing.T) {
	actual1 := EvaluateCriteria(111111)
	if !reflect.DeepEqual(actual1, true) {
		t.Errorf("Expected: %v, Actual: %v", true, actual1)
	}

	actual2 := EvaluateCriteria(223450)
	if !reflect.DeepEqual(actual2, false) {
		t.Errorf("Expected: %v, Actual: %v", false, actual2)
	}

	actual3 := EvaluateCriteria(123789)
	if !reflect.DeepEqual(actual3, false) {
		t.Errorf("Expected: %v, Actual: %v", false, actual3)
	}

	actual4 := EvaluateCriteria(111123)
	if !reflect.DeepEqual(actual4, true) {
		t.Errorf("Expected: %v, Actual: %v", true, actual4)
	}

	actual5 := EvaluateCriteria(135679)
	if !reflect.DeepEqual(actual5, false) {
		t.Errorf("Expected: %v, Actual: %v", false, actual5)
	}
}
