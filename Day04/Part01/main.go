package main

import (
	"fmt"
	"strconv"
)

func EvaluateCriteria(value int) bool {
	tempValue := strconv.Itoa(value)

	increase, sameState := false, false
	for i := 0; i < len(tempValue)-1; i++ {
		a, _ := strconv.Atoi(string(tempValue[i]))
		b, _ := strconv.Atoi(string(tempValue[i+1]))
		increase = a != b && a > b
		if !sameState && a == b {
			sameState = a == b
		}
		if increase {
			return false
		}
	}
	return sameState
}

func main() {
	counter := 0
	for i := 165432; i < 707912; i++ {
		if EvaluateCriteria(i) {
			counter++
		}
	}
	fmt.Println(counter)
}
