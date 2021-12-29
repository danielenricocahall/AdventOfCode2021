package main

import (
	. "AdventOfCode/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readLines(path string,
	computeCoordinateFunction func(x1 int, x2 int,
		y1 int, y2 int) []Coordinate) []Coordinate {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var coordinates []Coordinate
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")
		begin := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")
		x1, _ := strconv.Atoi(begin[0])
		y1, _ := strconv.Atoi(begin[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])
		coordinates = append(coordinates, computeCoordinateFunction(x1, x2, y1, y2)...)

	}
	return coordinates
}

func computeCoordinatesOnlyHorizontalAndVertical(
	x1 int,
	x2 int,
	y1 int,
	y2 int) []Coordinate {
	var coordinates []Coordinate

	if x1 == x2 {
		x := x1
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			coordinate := Coordinate{X: x, Y: y}
			coordinates = append(coordinates, coordinate)
		}

	} else if y1 == y2 {
		y := y1
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			coordinate := Coordinate{X: x, Y: y}
			coordinates = append(coordinates, coordinate)
		}
	}
	return coordinates
}

func computeCoordinatesAll(
	x1 int,
	x2 int,
	y1 int,
	y2 int) []Coordinate {
	var coordinates []Coordinate
	if x1 != x2 && y1 != y2 {

		for x, y := x1, y1; x != x2 && y != y2; {
			coordinate := Coordinate{X: x, Y: y}
			coordinates = append(coordinates, coordinate)
			if x1 > x2 {
				x--
			} else {
				x++
			}
			if y1 > y2 {
				y--
			} else {
				y++
			}
		}
		coordinates = append(coordinates, Coordinate{X: x2, Y: y2})
	} else {
		coordinates = append(coordinates,
			computeCoordinatesOnlyHorizontalAndVertical(x1, x2, y1, y2)...)
	}
	return coordinates
}
func countCoordinatesOccurrences(coordinates []Coordinate) map[Coordinate]int {
	occurrenceOfCoordinates := make(map[Coordinate]int)
	for _, coordinate := range coordinates {
		if _, ok := occurrenceOfCoordinates[coordinate]; ok {
			occurrenceOfCoordinates[coordinate] += 1
		} else {
			occurrenceOfCoordinates[coordinate] = 1
		}
	}
	return occurrenceOfCoordinates
}

func filterCoordinateOccurrences(
	coordinateOccurrences map[Coordinate]int,
	filterFunction func(coordinate Coordinate, occurrence int) bool) {
	for coordinate, occurrence := range coordinateOccurrences {
		if !filterFunction(coordinate, occurrence) {
			delete(coordinateOccurrences, coordinate)
		}
	}
}

func main() {
	coordinates := readLines("day_5/day_5.txt", computeCoordinatesAll)
	coordinateOccurrences := countCoordinatesOccurrences(coordinates)
	println(len(coordinateOccurrences))
	for k, v := range coordinateOccurrences {
		println("Key: (", k.X, k.Y, ") =>", v)
	}
	filterCoordinateOccurrences(coordinateOccurrences,
		func(coordinate Coordinate, occurrence int) bool {
			return occurrence > 1
		})
	println(len(coordinateOccurrences))

}
