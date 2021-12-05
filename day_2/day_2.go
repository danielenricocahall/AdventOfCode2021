package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	Down    Direction = "down"
	Up                = "up"
	Forward           = "forward"
)

type Vector struct {
	direction Direction
	magnitude int
}

func readLines(path string) ([]Vector, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var vectors []Vector
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		s := strings.Split(value, " ")
		magnitude, _ := strconv.Atoi(s[1])
		vector := Vector{Direction(s[0]),
			magnitude}
		vectors = append(vectors, vector)
	}
	return vectors, scanner.Err()
}

func computeTotalDistances(vectors []Vector) (int, int) {
	var totalVerticalDistance int
	var totalHorizontalDistance int
	for _, vector := range vectors {
		if vector.direction == Up {
			totalVerticalDistance -= vector.magnitude
		}
		if vector.direction == Down {
			totalVerticalDistance += vector.magnitude
		}
		if vector.direction == Forward {
			totalHorizontalDistance += vector.magnitude
		}
	}
	return totalHorizontalDistance, totalVerticalDistance
}

func computeTotalDistancesNewMethod(vectors []Vector) (int, int) {
	var aim int
	var totalHorizontalDistance int
	var totalVerticalDistance int
	for _, vector := range vectors {
		if vector.direction == Up {
			aim -= vector.magnitude
		}
		if vector.direction == Down {
			aim += vector.magnitude
		}
		if vector.direction == Forward {
			totalHorizontalDistance += vector.magnitude
			totalVerticalDistance += aim * vector.magnitude
		}
	}
	return totalHorizontalDistance, totalVerticalDistance
}
func main() {
	vectors, _ := readLines("day_2/day_2.txt")
	horizontal, vertical := computeTotalDistances(vectors)
	println(horizontal * vertical)
	newHorizontal, newVertical := computeTotalDistancesNewMethod(vectors)
	println(newHorizontal * newVertical)
}
