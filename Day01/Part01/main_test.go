package main

import (
	"testing"
)

func TestGetFuel(t *testing.T)  {
	fuelCase1 := GetFuel(12)
	if fuelCase1 != 2 {
		t.Errorf("Expected: %d, Actual: %d", 2, fuelCase1)
	}

	fuelCase2 := GetFuel(14)
	if fuelCase2 != 2 {
		t.Errorf("Expected: %d, Actual: %d", 2, fuelCase2)
	}

	fuelCase3 := GetFuel(1969)
	if fuelCase3 != 654 {
		t.Errorf("Expected: %d, Actual: %d", 654, fuelCase3)
	}

	fuelCase4 := GetFuel(100756)
	if fuelCase4 != 33583 {
		t.Errorf("Expected: %d, Actual: %d", 33583, fuelCase4)
	}
}