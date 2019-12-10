package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func GetInitialState(opcodes []int) []int {
	for i := 0; i < len(opcodes); i++ {
		var x, y, z, pos int

		if opcodes[i] == 99 {
			break
		}
		if opcodes[i] == 1 || opcodes[i] == 2 {
			x = opcodes[i+1]
			y = opcodes[i+2]
			pos = opcodes[i+3]
			if len(opcodes) > x && len(opcodes) > y {

				if opcodes[i] == 1 {
					z = opcodes[x] + opcodes[y]
				} else if opcodes[i] == 2 {
					z = opcodes[x] * opcodes[y]
				}
				opcodes[pos] = z
			}
			i += 3
		}
	}
	return opcodes
}

func main() {
	opcodes := []int{}
	data, err := ioutil.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	strOpcodes := strings.Split(string(data), ",")

	for _, i := range strOpcodes {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		opcodes = append(opcodes, j)
	}
	fmt.Println(GetInitialState(opcodes))
}
