package main

import (
	. "AdventOfCode/utils"
	"bufio"
	"fmt"
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

func findPathWithLowestRisk(cavern [][]int) []int {
	var results []int
	currentSum := 0
	memo := make(map[int]map[int][]int)
	computeNextStepWithLowestRisk(&cavern, 0, 0, currentSum, &memo, &results)
	//bs, _ := json.Marshal(memo)
	//fmt.Println(string(bs))
	return results
}

func computeNextStepWithLowestRisk(cavern *[][]int,
	horizontalPosition int,
	verticalPosition int,
	currentSum int,
	memo *map[int]map[int][]int,
	results *[]int) []int {
	rows := len(*cavern) - 1
	cols := len((*cavern)[0]) - 1
	if (*memo)[verticalPosition] == nil {
		println("Creating entry for " + strconv.Itoa(verticalPosition))
		(*memo)[verticalPosition] = make(map[int][]int)
	}
	if verticalPosition == rows && horizontalPosition == cols {
		(*memo)[verticalPosition][horizontalPosition] = append((*memo)[verticalPosition][horizontalPosition], []int{(*cavern)[verticalPosition][horizontalPosition] + currentSum}...)
		*results = append(*results, currentSum+(*cavern)[verticalPosition][horizontalPosition]-(*cavern)[0][0])
	} else if (*memo)[verticalPosition][horizontalPosition] != nil &&
		(currentSum+(*cavern)[verticalPosition][horizontalPosition] > (*memo)[verticalPosition][horizontalPosition][0] ||
			(len((*memo)[verticalPosition][horizontalPosition]) == 2 &&
				currentSum+(*cavern)[verticalPosition][horizontalPosition] > (*memo)[verticalPosition][horizontalPosition][1])) {
		println("Hit cache for (" + strconv.Itoa(verticalPosition) + "," + strconv.Itoa(horizontalPosition) + ")")
		println(currentSum + (*cavern)[verticalPosition][horizontalPosition])
		//fmt.Println((*memo)[verticalPosition][horizontalPosition])
		return (*memo)[verticalPosition][horizontalPosition]

	} else {
		if verticalPosition+1 <= rows {
			(*memo)[verticalPosition][horizontalPosition] = append((*memo)[verticalPosition][horizontalPosition],
				computeNextStepWithLowestRisk(cavern,
					horizontalPosition, verticalPosition+1, currentSum+(*cavern)[verticalPosition][horizontalPosition], memo, results)...)
		}
		if horizontalPosition+1 <= cols {
			(*memo)[verticalPosition][horizontalPosition] = append((*memo)[verticalPosition][horizontalPosition],
				computeNextStepWithLowestRisk(cavern,
					horizontalPosition+1, verticalPosition, currentSum+(*cavern)[verticalPosition][horizontalPosition], memo, results)...)
		}
	}
	return []int{(*cavern)[verticalPosition][horizontalPosition] + currentSum}

}

func main() {
	cavern := readLines("day_15/test.txt")
	fmt.Println(cavern)
	results := findPathWithLowestRisk(cavern)
	sort.Ints(results)
	//fmt.Println(results)
	println(len(results))
	println(results[0])
}
