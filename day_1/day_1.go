package main

import (
	"bufio"
	"os"
	"strconv"
)

func findBiggestDepthDifference(depths []int) int {
	var i int
	var prevDepth = depths[0]
	var numberOfTimesDepthIncreased = 0
	for i = 1; i < len(depths); i++ {
		if depths[i] > prevDepth {
			numberOfTimesDepthIncreased += 1
		}
		prevDepth = depths[i]
	}
	return numberOfTimesDepthIncreased
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}
	return lines, scanner.Err()
}

func main() {
	dat, _ := readLines("day_1/day_1.txt")
	println(findBiggestDepthDifference(dat))
	var i, j int
	var slidingWindowSums []int
	for j = 0; j < len(dat) - 2; j++ {
		var currentSum = 0
		for i = 0; i < 3; i++ {
			currentSum += dat[j + i]
		}
		slidingWindowSums = append(slidingWindowSums, currentSum)
	}
	println(findBiggestDepthDifference(slidingWindowSums))
}
