package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LowPoint struct {
	row   int
	col   int
	value int
}

func readLines(path string) [][]int {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var results [][]int
	for scanner.Scan() {
		results = append(results, utils.Map(strings.Split(scanner.Text(), ""), strconv.Atoi))
	}
	return results
}

func findLowPoints(heights [][]int) []LowPoint {
	var lowPoints []LowPoint
	for row, _ := range heights {
		for col, _ := range heights[row] {
			condition := true
			if row > 0 {
				condition = condition && (heights[row][col] < heights[row-1][col])
			}
			if row < len(heights)-1 {
				condition = condition && (heights[row][col] < heights[row+1][col])
			}
			if col > 0 {
				condition = condition && (heights[row][col] < heights[row][col-1])
			}
			if col < len(heights[row])-1 {
				condition = condition && (heights[row][col] < heights[row][col+1])
			}
			if condition {
				lowPoints = append(lowPoints, LowPoint{
					row:   row,
					col:   col,
					value: heights[row][col],
				})
			}
		}
	}
	return lowPoints
}

func computeTotalRiskLevel(lowPoints []LowPoint) int {
	totalRiskLevel := 0
	for _, lowPoint := range lowPoints {
		totalRiskLevel += lowPoint.value + 1
	}
	return totalRiskLevel
}

func findBasin(heights *[][]int, visitedTracking *[][]bool, row int, col int) []int {
	var results []int
	if row < 0 || col < 0 || row > len(*heights)-1 || col > len((*heights)[row])-1 {
		return results
	}
	height := (*heights)[row][col]
	visited := (*visitedTracking)[row][col]
	if height == 9 || visited {
		return results
	}
	fmt.Printf("(%d, %v): %v\n", row, col, height)
	results = append(results, height)
	(*visitedTracking)[row][col] = true
	if row > 0 {
		results = append(results, findBasin(heights, visitedTracking, row-1, col)...)
	}
	if col > 0 {
		results = append(results, findBasin(heights, visitedTracking, row, col-1)...)
	}
	if row < len(*heights)-1 {
		results = append(results, findBasin(heights, visitedTracking, row+1, col)...)
	}
	if col < len((*heights)[row])-1 {
		results = append(results, findBasin(heights, visitedTracking, row, col+1)...)
	}
	return results
}

func createVisitedTrackingSlice(heights *[][]int) [][]bool {
	visited := make([][]bool, len(*heights))
	for row, _ := range *heights {
		visited[row] = make([]bool, len((*heights)[row]))
		for col, _ := range (*heights)[row] {
			visited[row][col] = false
		}
	}
	return visited
}

func findBasins(heights *[][]int, lowPoints []LowPoint) [][]int {
	var basins [][]int
	visited := createVisitedTrackingSlice(heights)
	for _, lowPoint := range lowPoints {
		fmt.Printf("Low Point: (%d, %v) %v\n", lowPoint.row, lowPoint.col, lowPoint.value)
		println("Basin: ")
		basin := findBasin(heights, &visited, lowPoint.row, lowPoint.col)
		basins = append(basins, basin)
		fmt.Printf("Size of basin: %v\n", len(basin))
	}
	return basins
}

func main() {
	heights := readLines("day_9/data.txt")
	lowPoints := findLowPoints(heights)
	totalRisk := computeTotalRiskLevel(lowPoints)
	println(totalRisk)
	basins := findBasins(&heights, lowPoints)
	sort.Slice(basins, func(i int, j int) bool {
		return len(basins[i]) < len(basins[j])
	})
	println(len(basins[len(basins)-1]) *
		len(basins[len(basins)-2]) *
		len(basins[len(basins)-3]))

}
