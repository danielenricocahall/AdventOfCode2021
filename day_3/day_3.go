package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var bits []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bitString := scanner.Text()
		bits = append(bits, bitString)
	}
	return bits, scanner.Err()
}

func determineMostCommonBitPerColumn(rowsOfBits []string) string {
	var sumPerColumn = computeSumByColumn(rowsOfBits)
	return constructBitString(sumPerColumn, rowsOfBits)
}

func computeSumByColumn(rowsOfBits []string) []int {
	var j int
	var sumPerColumn = make([]int, len(rowsOfBits[0]))
	for _, bitString := range rowsOfBits {
		for _, bitChar := range bitString {
			bit, _ := strconv.Atoi(string(bitChar))
			sumPerColumn[j] += bit
		}
	}
	return sumPerColumn
}

func constructBitString(sumPerColumn []int, rowsOfBits []string) string {
	var result string
	for _, sum := range sumPerColumn {
		if sum >= len(rowsOfBits)/2 {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func flipBits(bitString string) string {
	var invertedBitString = ""
	for bit := range bitString {
		if string(rune(bit)) == "1" {
			invertedBitString += "0"
		} else {
			invertedBitString += "1"
		}
	}
	return invertedBitString
}

func computeDiagnostics(bits []string, compareFunction func(x int, y int) bool) []string {
	var i int
	for i = 0; i < len(bits[0]); i++ {
		var sumForColumn = 0
		var j int
		for j = 0; j < len(bits); j++ {
			bit, _ := strconv.Atoi(string(bits[j][i]))
			sumForColumn += bit
		}
		var filterFunction func(s string) bool
		var majorityOfRows = int(math.Ceil(float64(len(bits)) / 2.0))
		if compareFunction(sumForColumn, majorityOfRows) {
			filterFunction = func(s string) bool {
				return string(s[i]) == "1"
			}
		} else {
			filterFunction = func(s string) bool {
				return string(s[i]) == "0"
			}
		}
		bits = filter(bits, filterFunction)
		if len(bits) == 1 {
			break
		}
	}
	return bits
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func main() {
	dat, _ := readLines("day_3/day_3_test.txt")
	mostCommonBitsPerColumn := determineMostCommonBitPerColumn(dat)
	leastCommonBitsPerColumn := flipBits(mostCommonBitsPerColumn)
	epsilonRate, _ := strconv.ParseInt(mostCommonBitsPerColumn, 2, 64)
	gammaRate, _ := strconv.ParseInt(leastCommonBitsPerColumn, 2, 64)
	println(mostCommonBitsPerColumn)
	println(leastCommonBitsPerColumn)
	println(epsilonRate)
	println(gammaRate)
	println(epsilonRate * gammaRate)
	var oxygenDiagnosticCompare = func(x int, y int) bool { return x >= y }
	var co2ScrubberCompare = func(x int, y int) bool { return x < y }
	oxygenDiagnostic := computeDiagnostics(dat, oxygenDiagnosticCompare)[0]
	co2ScrubberRating := computeDiagnostics(dat, co2ScrubberCompare)[0]
	foo, _ := strconv.ParseInt(oxygenDiagnostic, 2, 64)
	bar, _ := strconv.ParseInt(co2ScrubberRating, 2, 64)
	println(foo)
	println(bar)
	println(foo * bar)
}
