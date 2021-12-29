package main

import (
	. "AdventOfCode/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Fold struct {
	axis  string
	value int
}

func readLines(path string) ([]Coordinate, []Fold) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var folds []Fold
	var results []Coordinate
	for scanner.Scan() {
		text := scanner.Text()
		println(text)
		if strings.Contains(text, "fold") {
			splitted := strings.Split(text, "fold along ")
			fold := strings.Split(splitted[1], "=")
			value, _ := strconv.Atoi(fold[1])
			folds = append(folds, Fold{axis: fold[0], value: value})
		}
		if !strings.Contains(text, "fold") && len(text) > 0 {
			row := Map(strings.Split(text, ","), strconv.Atoi)
			results = append(results, Coordinate{X: row[0], Y: row[1]})
		}
	}
	return results, folds
}

func getMaxX(coordinates *[]Coordinate) int {
	max := 0
	for _, coordinate := range *coordinates {
		if coordinate.X > max {
			max = coordinate.X
		}
	}
	return max
}

func getMaxY(coordinates *[]Coordinate) int {
	max := 0
	for _, coordinate := range *coordinates {
		if coordinate.Y > max {
			max = coordinate.Y
		}
	}
	return max
}

func createPaper(coordinates *[]Coordinate) [][]string {
	maxX := getMaxX(coordinates) + 1
	maxY := getMaxY(coordinates) + 1
	paper := make([][]string, maxY)
	for y := 0; y < maxY; y++ {
		paper[y] = make([]string, maxX)
		for x := 0; x < maxX; x++ {
			paper[y][x] = "."
		}
	}
	for _, coordinate := range *coordinates {
		paper[coordinate.Y][coordinate.X] = "#"
	}
	return paper
}

func printPaper(paper *[][]string) {
	println()
	for _, row := range *paper {
		for col := 0; col < len(row); col++ {
			print(row[col])
		}
		println()
	}
}

func addFoldingLine(paper *[][]string, fold Fold) {
	if fold.axis == "x" {
		for i, _ := range *paper {
			(*paper)[i][fold.value] = "|"
		}
	} else if fold.axis == "y" {
		for i, _ := range (*paper)[0] {
			(*paper)[fold.value][i] = "-"
		}
	}
	printPaper(paper)

}

func foldUp(paper *[][]string, value int) {
	maxY := len(*paper) - 1
	for row, _ := range (*paper)[:value] {
		for col, _ := range (*paper)[row] {
			if (*paper)[maxY-row][col] == "#" {
				(*paper)[row][col] = (*paper)[maxY-row][col]
			}
		}
	}
	*paper = (*paper)[:value]
}

func foldLeft(paper *[][]string, value int) {
	maxX := len((*paper)[0]) - 1
	for row, _ := range *paper {
		for col, _ := range (*paper)[row][:value] {
			if (*paper)[row][col] != "#" {
				(*paper)[row][col] = (*paper)[row][maxX-col]
			}
		}
	}
	for row, _ := range *paper {
		(*paper)[row] = (*paper)[row][:value]
	}
}

func foldPaper(paper *[][]string, folds []Fold, numFolds int) {
	folds = folds[:numFolds]
	for _, fold := range folds {
		if fold.axis == "y" {
			foldUp(paper, fold.value)
			//printPaper(paper)
		} else if fold.axis == "x" {
			foldLeft(paper, fold.value)
			//printPaper(paper)
		}
	}
}

func countVisibleDots(paper *[][]string) int {
	count := 0
	for row, _ := range *paper {
		for _, dot := range (*paper)[row] {
			if dot == "#" {
				count += 1
			}
		}
	}
	return count
}

func replaceNotVisibleDots(paper *[][]string) {
	for row, _ := range *paper {
		for col, dot := range (*paper)[row] {
			if dot == "." {
				(*paper)[row][col] = " "
			}
		}
	}
}

func main() {
	/*

		####..##..#..#.#..#.###..####..##..###.
		#....#..#.#..#.#.#..#..#.#....#..#.#..#
		###..#..#.####.##...#..#.###..#....#..#
		#....####.#..#.#.#..###..#....#....###.
		#....#..#.#..#.#.#..#.#..#....#..#.#...
		####.#..#.#..#.#..#.#..#.####..##..#...


		####.####.#..#.#.##.#.##.####..#.#.####.
		###..####.####.###..##...####.##...####.
		###..####.####.###..####.###..#..#.##.#.
		###..####.####.###..####.###..#....####.
		####.####.#..#.#.##.####.###..#.##.#....
		####.#..#.#..#.#.##.#.##.###..####.#....


	*/
	coordinates, folds := readLines("day_13/data.txt")
	paper := createPaper(&coordinates)
	foldPaper(&paper, folds, len(folds))

	//println(countVisibleDots(&paper))
	printPaper(&paper)

}
