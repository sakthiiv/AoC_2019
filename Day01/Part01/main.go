package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func GetFuel(mass int) int64 {
	x := int64(mass/3) - 2
	return x
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
			totalFuel += GetFuel(module)
		}

	}
	fmt.Println(totalFuel)
}
