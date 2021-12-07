package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func readLines(path string) []Coordinate {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var coordinates []Coordinate
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")
		begin := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")
		x1, y1 := begin[0], begin[1]
		x2, y2 := end[0], end[1]
		if x1 == x2 {
			x, _ := strconv.Atoi(x1)
			y1, _ := strconv.Atoi(y1)
			y2, _ := strconv.Atoi(y2)
			for y := y1; y < y2; y++ {
				coordinate := Coordinate{x: x, y: y}
				coordinates = append(coordinates, coordinate)
			}

		} else if y1 == y2 {
			x1, _ := strconv.Atoi(x1)
			x2, _ := strconv.Atoi(x2)
			y, _ := strconv.Atoi(y1)
			for x := x1; x < x2; x++ {
				coordinate := Coordinate{x: x, y: y}
				coordinates = append(coordinates, coordinate)
			}
		}

	}
	return coordinates
}

func main() {
	println(readLines("day_5/day_5_test.txt"))
}
