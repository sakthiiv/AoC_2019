package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func GetFuel(mass int64) int64 {
	x := int64(mass/3) - 2
	return x
}

func CalFuelForFuel(fuel int64, sum int64) int64 {
	fuelForFuel := GetFuel(fuel)
	if fuelForFuel <= 0 {
		return sum
	}
	sum += fuelForFuel
	return CalFuelForFuel(fuelForFuel, sum)
}

func main() {

	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	modules := strings.Split(string(data), "\n")
	var totalFuel int64 = 0
	for i := 0; i < len(modules); i++ {
		if modules[i] != "" {
			module, err := strconv.Atoi(modules[i])
			if err != nil {
				fmt.Println(err)
			}
			fuel := GetFuel(int64(module))
			totalFuel += fuel
			totalFuel += CalFuelForFuel(fuel, 0)
		}
	}
	fmt.Println(totalFuel)
}
