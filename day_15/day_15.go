package main

import (
	. "AdventOfCode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func readLines(path string) [][]int {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var results [][]int
	for scanner.Scan() {
		row := Map(strings.Split(scanner.Text(), ""), strconv.Atoi)
		results = append(results, row)
	}
	return results
}

func expandHorizontal(cavern *[][]int, n int) {
	for i, row := range *cavern {
		for j := 0; j < n-1; j++ {
			for _, riskLevel := range row {
				(*cavern)[i] = append((*cavern)[i], ((riskLevel+j)%9)+1)
			}
		}
	}
}

func expandVertical(cavern *[][]int, n int) {
	var newRows [][]int
	for j := 0; j < n-1; j++ {
		for _, row := range *cavern {
			var newRow []int
			for _, riskLevel := range row {
				newRow = append(newRow, ((riskLevel+j)%9)+1)
			}
			newRows = append(newRows, newRow)
		}
	}
	*cavern = append(*cavern, newRows...)
}

func expandCavern(cavern *[][]int, n int) {
	expandHorizontal(cavern, n)
	expandVertical(cavern, n)

}

func createMemoTable(cavern *[][]int) [][]int {
	memo := make([][]int, len(*cavern))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(memo))
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = math.MaxInt
		}
	}
	return memo

}

func findPathWithLowestRisk(cavern [][]int) int {
	minTotalRisk := math.MaxInt
	currentSum := 0
	memo := createMemoTable(&cavern)
	computeNextStepWithLowestRisk(&cavern, 0, 0, currentSum, &memo, &minTotalRisk)
	//bs, _ := json.Marshal(memo)
	//fmt.Println(string(bs))
	return minTotalRisk
}

func computeNextStepWithLowestRisk(cavern *[][]int,
	horizontalPosition int,
	verticalPosition int,
	currentSum int,
	memo *[][]int,
	minTotalRisk *int) {
	rows := len(*cavern) - 1
	cols := len((*cavern)[0]) - 1

	if verticalPosition == rows && horizontalPosition == cols {
		totalRisk := currentSum + (*cavern)[verticalPosition][horizontalPosition]
		if totalRisk < *minTotalRisk {
			*minTotalRisk = totalRisk
		}
	} else {
		currentSum += (*cavern)[verticalPosition][horizontalPosition]
		(*memo)[verticalPosition][horizontalPosition] = currentSum
		if verticalPosition+1 <= rows &&
			(currentSum+(*cavern)[verticalPosition+1][horizontalPosition] < (*memo)[verticalPosition+1][horizontalPosition]) &&
			*minTotalRisk > (currentSum+(*cavern)[verticalPosition+1][horizontalPosition]) {
			computeNextStepWithLowestRisk(cavern,
				horizontalPosition, verticalPosition+1, currentSum, memo, minTotalRisk)
		}
		if horizontalPosition+1 <= cols &&
			(currentSum+(*cavern)[verticalPosition][horizontalPosition+1] < (*memo)[verticalPosition][horizontalPosition+1]) &&
			*minTotalRisk > currentSum+(*cavern)[verticalPosition][horizontalPosition+1] {
			computeNextStepWithLowestRisk(cavern,
				horizontalPosition+1, verticalPosition, currentSum, memo, minTotalRisk)
		}
		if verticalPosition-1 >= 0 &&
			(currentSum+(*cavern)[verticalPosition-1][horizontalPosition] < (*memo)[verticalPosition-1][horizontalPosition]) &&
			*minTotalRisk > currentSum+(*cavern)[verticalPosition-1][horizontalPosition] {
			computeNextStepWithLowestRisk(cavern,
				horizontalPosition, verticalPosition-1, currentSum, memo, minTotalRisk)
		}
		if horizontalPosition-1 >= 0 &&
			(currentSum+(*cavern)[verticalPosition][horizontalPosition-1] < (*memo)[verticalPosition][horizontalPosition-1] &&
				*minTotalRisk > currentSum+(*cavern)[verticalPosition][horizontalPosition-1]) {
			computeNextStepWithLowestRisk(cavern,
				horizontalPosition-1, verticalPosition, currentSum, memo, minTotalRisk)
		}
	}

}

func main() {
	cavern := readLines("day_15/data.txt")
	expandCavern(&cavern, 5)
	//fmt.Println(cavern)
	start := time.Now()
	path := findPathWithLowestRisk(cavern)
	duration := time.Since(start)
	fmt.Println(duration)
	//fmt.Println(results)
	//fmt.Println(results)
	println(path - cavern[0][0])
}
