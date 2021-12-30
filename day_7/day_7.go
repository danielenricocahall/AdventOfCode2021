package main

import (
	. "AdventOfCode/utils"
	"math"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) []int {
	data, _ := os.ReadFile(path)
	var positions []int
	for _, value := range strings.Split(string(data), ",") {
		position, _ := strconv.Atoi(value)
		positions = append(positions, position)
	}
	return positions
}

func computeL1(x int, k int) int {
	return int(math.Abs(float64(x - k)))
}

func computeSumOfDifferences(positions []int, k int) int {
	acc := 0
	for _, position := range positions {
		acc += computeL1(position, k)
	}
	return acc
}

func computeSumOfTriangularSumDifferences(positions []int, k int) int {
	acc := 0
	for _, position := range positions {
		n := computeL1(position, k)
		acc += n * (n + 1) / 2
	}
	return acc
}

func computeMinimumFuel(positions []int,
	differenceFunction func([]int, int) int) int {
	maxPosition := ComputeMax(positions)
	minFuelUsed := math.Inf(1)
	for k := 0; k <= maxPosition; k++ {
		fuelUsed := differenceFunction(positions, k)
		if float64(fuelUsed) < minFuelUsed {
			minFuelUsed = float64(fuelUsed)
		}
	}
	return int(minFuelUsed)
}

func main() {
	positions := readLines("day_7/data.txt")
	println(computeMinimumFuel(positions, computeSumOfTriangularSumDifferences))
}
