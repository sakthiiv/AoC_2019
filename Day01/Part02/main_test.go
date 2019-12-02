package main

import (
	"fmt"
	"testing"
)

func TestGetFuel(t *testing.T) {
	fmt.Println("Calcuate fuel")

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

func TestCalFuelForFuel(t *testing.T) {
	fmt.Println("Calcuate fuel for fuel")

	fuelForFuelCase1 := CalFuelForFuel(14, 0)
	if fuelForFuelCase1 != 2 {
		t.Errorf("Expected: %d, Actual: %d", 2, fuelForFuelCase1)
	}

	fuelForFuelCase2 := CalFuelForFuel(1969, 0)
	if fuelForFuelCase2 != 966 {
		t.Errorf("Expected: %d, Actual: %d", 966, fuelForFuelCase2)
	}

	fuelForFuelCase3 := CalFuelForFuel(100756, 0)
	if fuelForFuelCase3 != 50346 {
		t.Errorf("Expected: %d, Actual: %d", 50346, fuelForFuelCase3)
	}
}
