package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func add(mode1, mode2, idx int, data []int) int {
	x, y := getParams(mode1, data, idx+1), getParams(mode2, data, idx+2)
	return x + y
}

func multiply(mode1, mode2, idx int, data []int) int {
	x, y := getParams(mode1, data, idx+1), getParams(mode2, data, idx+2)
	return x * y
}

func output(o int) {
	fmt.Println("Output code: ", o)
}

func getModes(x int) []int {
	modes := []int{}
	paddedModes := fmt.Sprintf("%05d", x)[0:3]
	for _, v := range reverse(paddedModes) {
		mode, _ := strconv.Atoi(string(v))
		modes = append(modes, mode)
	}
	return modes
}

func getParams(mode int, data []int, idx int) int {
	if mode == 0 {
		return data[data[idx]]
	} else if mode == 1 {
		return data[idx]
	}
	return 0
}

func reverse(s string) []rune {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func PerformDiagnosis(data []int, input int) []int {
	idx, result, code := 0, 0, 0
	for data[idx] != 99 {
		opcode := data[idx] % 100
		modes := getModes(data[idx])
		mode1, mode2 := modes[0], modes[1]
		switch opcode {
		case 1:
			result = add(mode1, mode2, idx, data)
			data[data[idx+3]] = result
			idx += 4
		case 2:
			result = multiply(mode1, mode2, idx, data)
			data[data[idx+3]] = result
			idx += 4
		case 3:
			data[data[idx+1]] = input
			idx += 2
		case 4:
			code = data[data[idx+1]]
			idx += 2
		}
	}
	fmt.Println(code)
	return data
}

func main() {
	instructions := []int{}
	data, _ := ioutil.ReadFile("../input")

	strOpcodes := strings.Split(string(data), ",")

	for _, v := range strOpcodes {
		i, _ := strconv.Atoi(v)
		instructions = append(instructions, i)
	}
	PerformDiagnosis(instructions, 1)
}
