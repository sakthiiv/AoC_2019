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

func EvaluateLargerGroup(value int) bool {
	tempValue := strconv.Itoa(value)
	state := false
	valueMap := make(map[string]int)
	for i := 0; i < len(tempValue); i++ {
		valueMap[string(tempValue[i])] = valueMap[string(tempValue[i])] + 1
	}

	for _, v := range valueMap {
		if !state && v == 2 {
			state = true
		}
	}
	return state
}

func Evaluate(value int) bool {
	return EvaluateLargerGroup(value) && EvaluateCriteria(value)
}

func main() {
	counter := 0
	for i := 165432; i < 707912; i++ {
		if Evaluate(i) {
			counter++
		}
	}
	fmt.Println(counter)
}
