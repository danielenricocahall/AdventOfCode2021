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

func increaseEnergyLevels(
	octopi *[][]int) {
	rows := len(*octopi)
	cols := len((*octopi)[0])
	flash(octopi,
		0,
		rows,
		0,
		cols)
}

func flash(octopi *[][]int,
	rowStart int,
	rowEnd int,
	colStart int,
	colEnd int) {
	rowStart = utils.Max(rowStart, 0)
	rowEnd = utils.Min(rowEnd, len(*octopi)-1)
	colStart = utils.Max(colStart, 0)
	colEnd = utils.Min(colEnd, len((*octopi)[0])-1)
	for row := rowStart; row <= rowEnd; row++ {
		for col := colStart; col <= colEnd; col++ {
			(*octopi)[row][col] += 1
		}
	}
}

func evaluateEnergyLevels(
	octopi *[][]int,
	flashed *[][]bool,
	rowStart int,
	rowEnd int,
	colStart int,
	colEnd int) {
	rowStart = utils.Max(rowStart, 0)
	rowEnd = utils.Min(rowEnd, len(*octopi)-1)
	colStart = utils.Max(colStart, 0)
	colEnd = utils.Min(colEnd, len((*octopi)[0])-1)
	for row := rowStart; row <= rowEnd; row++ {
		for col := colStart; col <= colEnd; col++ {
			energyLevel := (*octopi)[row][col]
			if energyLevel > 9 && !(*flashed)[row][col] {
				(*flashed)[row][col] = true
				flash(octopi, row-1, row+1, col-1, col+1)
				evaluateEnergyLevels(octopi, flashed, row-1, row+1, col-1, col+1)
			}
		}
	}
}

func printOctopiEnergyStates(octopi *[][]int, step int) {
	println("STEP " + strconv.Itoa(step+1))
	for _, row := range *octopi {
		for _, col := range row {
			print(col)
		}
		println()
	}
	println()
}

func createFlashedSlice(octopi *[][]int) [][]bool {
	flashed := make([][]bool, len(*octopi))
	for row, _ := range *octopi {
		flashed[row] = make([]bool, len((*octopi)[row]))
		for col, _ := range (*octopi)[row] {
			flashed[row][col] = false
		}
	}
	return flashed
}

func updateFlashedOctopi(octopi *[][]int, flashed *[][]bool) {
	for i, _ := range *flashed {
		for j, _ := range (*flashed)[i] {
			if (*flashed)[i][j] == true {
				(*octopi)[i][j] = 0
			}
		}
	}
}

func computeNumberOfFlashes(flashedStates *[][]bool) int {
	numberOfFlashes := 0
	for _, row := range *flashedStates {
		for _, flashed := range row {
			if flashed {
				numberOfFlashes += 1
			}
		}
	}
	return numberOfFlashes
}

func flashesSynchronized(flashedStates *[][]bool) bool {
	for _, row := range *flashedStates {
		for _, flashed := range row {
			if !flashed {
				return false
			}
		}
	}
	return true
}

func run(octopi *[][]int,
	steps int) {
	rows := len(*octopi)
	cols := len((*octopi)[0])
	numberOfFlashes := 0
	for step := 0; step < steps; step++ {
		flashed := createFlashedSlice(octopi)
		increaseEnergyLevels(octopi)
		evaluateEnergyLevels(octopi, &flashed, 0, rows, 0, cols)
		updateFlashedOctopi(octopi, &flashed)
		numberOfFlashes += computeNumberOfFlashes(&flashed)
		if flashesSynchronized(&flashed) {
			println("Flashes Synchronized on step " + strconv.Itoa(step+1))
		}
		//printOctopiEnergyStates(octopi, step)
	}
	println("Number of Flashes: " + strconv.Itoa(numberOfFlashes))
}

func main() {
	octopi := readLines("day_11/data.txt")
	run(&octopi, 400)
}
