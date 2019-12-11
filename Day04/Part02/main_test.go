package main

import (
	"reflect"
	"testing"
)

func TestEvaluateCriteria(t *testing.T) {

	actual2 := Evaluate(223450)
	if !reflect.DeepEqual(actual2, false) {
		t.Errorf("Expected: %v, Actual: %v", false, actual2)
	}

	actual3 := Evaluate(123789)
	if !reflect.DeepEqual(actual3, false) {
		t.Errorf("Expected: %v, Actual: %v", false, actual3)
	}

	actual5 := Evaluate(135679)
	if !reflect.DeepEqual(actual5, false) {
		t.Errorf("Expected: %v, Actual: %v", false, actual5)
	}

	actual6 := Evaluate(112233)
	if !reflect.DeepEqual(actual6, true) {
		t.Errorf("Expected: %v, Actual: %v", true, actual6)
	}

	actual7 := Evaluate(123444)
	if !reflect.DeepEqual(actual7, false) {
		t.Errorf("Expected: %v, Actual: %v", false, actual7)
	}

	actual8 := Evaluate(111122)
	if !reflect.DeepEqual(actual8, true) {
		t.Errorf("Expected: %v, Actual: %v", true, actual8)
	}

	actual9 := Evaluate(444556)
	if !reflect.DeepEqual(actual9, true) {
		t.Errorf("Expected: %v, Actual: %v", true, actual9)
	}
}
