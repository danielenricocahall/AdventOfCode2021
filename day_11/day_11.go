package main

import (
	"AdventOfCode/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) [][]int {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var results [][]int
	for scanner.Scan() {
		row := utils.Map(strings.Split(scanner.Text(), ""), strconv.Atoi)
		results = append(results, row)
	}
	return results
}

func increaseEnergyLevels(octopi *[][]int) {
	for i, row := range *octopi {
		for j, _ := range row {
			(*octopi)[i][j] += 1
		}
	}
}

func evaluateEnergyLevels(octopi [][]int) {
	for _, row := range octopi {
		for _, energyLevel := range row {
			if energyLevel > 9 {

			}
		}
	}
}

func run(octopi *[][]int, steps int) {
	for step := 0; step < steps; step++ {
		increaseEnergyLevels(octopi)
	}
}

func main() {
	octopi := readLines("day_11/test.txt")
	run(&octopi, 1)
}
