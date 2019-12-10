package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type wire struct {
	direction rune
	length    int
	currentX  int
	currentY  int
}

type point struct {
	x, y, distance int
}

func fetchInput() []string {
	data, err := ioutil.ReadFile("../input")
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(data), "\n")
}

func parseInputToPaths(inputs []string) [][]wire {

	ilen := len(inputs)
	wires := make([][]wire, ilen)

	for i := 0; i < ilen; i++ {
		directions := strings.Split(inputs[i], ",")
		dlen := len(directions)
		wires[i] = make([]wire, dlen)
		for j := 0; j < dlen; j++ {
			wirelen, _ := strconv.Atoi(directions[j][1:])
			wires[i][j] = wire{direction: (rune)(directions[j][0]), length: wirelen}
		}
	}

	return wires
}

func getSteps(paths []wire) []point {
	x, y, distance := 0, 0, 0
	points := make([]point, 0)
	for j := 0; j < len(paths); j++ {
		for l := 0; l < paths[j].length; l++ {
			distance += 1
			switch paths[j].direction {
			case 'R':
				x += 1
			case 'L':
				x -= 1
			case 'U':
				y -= 1
			case 'D':
				y += 1
			}
			p := point{x: x, y: y, distance: distance}
			points = append(points, p)
		}
	}
	return points
}

func main() {
	wires := parseInputToPaths(fetchInput())
	firstPath := getSteps(wires[0])
	secondPath := getSteps(wires[1])
	totalDistance := math.MaxInt32

	for k := 0; k < len(firstPath); k++ {
		for l := 0; l < len(secondPath); l++ {
			if firstPath[k].x == secondPath[l].x && firstPath[k].y == secondPath[l].y {
				currentDistance := firstPath[k].distance + secondPath[l].distance
				if totalDistance > currentDistance {
					totalDistance = currentDistance
				}
			}
		}
	}

	fmt.Println(totalDistance)
}
