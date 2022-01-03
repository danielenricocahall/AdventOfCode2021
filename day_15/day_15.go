package main

import (
	. "AdventOfCode/utils"
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
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

func findPathWithLowestRisk(cavern [][]int) []int {
	var results []int
	currentSum := 0
	memo := make([][]int, len(cavern))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(memo))
	}
	computeNextStepWithLowestRisk(&cavern, 0, 0, currentSum, &memo, &results)
	//bs, _ := json.Marshal(memo)
	//fmt.Println(string(bs))
	return results
}

func computeNextStepWithLowestRisk(cavern *[][]int,
	horizontalPosition int,
	verticalPosition int,
	currentSum int,
	memo *[][]int,
	results *[]int) int {
	rows := len(*cavern) - 1
	cols := len((*cavern)[0]) - 1

	if verticalPosition == rows && horizontalPosition == cols {
		*results = append(*results, currentSum+(*cavern)[verticalPosition][horizontalPosition])
	} else if (*memo)[verticalPosition][horizontalPosition] != 0 && currentSum+(*cavern)[verticalPosition][horizontalPosition] >= (*memo)[verticalPosition][horizontalPosition] {
		return (*memo)[verticalPosition][horizontalPosition]
	} else {
		(*memo)[verticalPosition][horizontalPosition] = (*cavern)[verticalPosition][horizontalPosition] + currentSum
		if verticalPosition+1 <= rows {
			(*memo)[verticalPosition+1][horizontalPosition] =
				computeNextStepWithLowestRisk(cavern,
					horizontalPosition, verticalPosition+1, currentSum+(*cavern)[verticalPosition][horizontalPosition], memo, results)
		}
		if horizontalPosition+1 <= cols {
			(*memo)[verticalPosition][horizontalPosition+1] =
				computeNextStepWithLowestRisk(cavern,
					horizontalPosition+1, verticalPosition, currentSum+(*cavern)[verticalPosition][horizontalPosition], memo, results)
		}
		if verticalPosition-1 >= 0 {
			(*memo)[verticalPosition-1][horizontalPosition] =
				computeNextStepWithLowestRisk(cavern,
					horizontalPosition, verticalPosition-1, currentSum+(*cavern)[verticalPosition][horizontalPosition], memo, results)
		}
		if horizontalPosition-1 >= 0 {
			(*memo)[verticalPosition][horizontalPosition-1] =
				computeNextStepWithLowestRisk(cavern,
					horizontalPosition-1, verticalPosition, currentSum+(*cavern)[verticalPosition][horizontalPosition], memo, results)
		}
	}
	return (*cavern)[verticalPosition][horizontalPosition] + currentSum

}

func main() {
	cavern := readLines("day_15/data.txt")
	expandCavern(&cavern, 5)
	//fmt.Println(cavern)
	results := findPathWithLowestRisk(cavern)
	sort.Ints(results)
	//fmt.Println(results)
	println(len(results))
	//fmt.Println(results)
	println(results[0] - cavern[0][0])
}
