package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type BingoCard = [5][]int

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
func Map(vs []string, f func(string) (int, error)) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i], _ = f(v)
	}
	return vsm
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func constructBingoCard(lines []string, i int) BingoCard {
	var bingoCard = BingoCard{}
	for j := 0; j < 5; j++ {
		line := strings.Split(lines[i+j], " ")
		line = filter(line, func(s string) bool {
			return len(s) > 0
		})
		bingoCard[j] = Map(line, strconv.Atoi)
	}
	return bingoCard
}

func sum(arr []int) int {
	acc := 0
	for _, v := range arr {
		acc += v
	}
	return acc
}
func checkIfWinner(bingoCard BingoCard, i int, j int) bool {
	n := len(bingoCard[0])
	var vertical []int
	horizontal := bingoCard[i]
	for k := 0; k < n; k++ {
		vertical = append(vertical, bingoCard[k][j])
	}
	return sum(horizontal) == -n || sum(vertical) == -n
}

func sumOfUnmarkedNumbers(bingoCard *BingoCard) int {
	n := len(bingoCard[0])
	acc := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if bingoCard[i][j] != -1 {
				acc += bingoCard[i][j]
			}
		}
	}
	return acc
}

func checkCard(bingoCard *BingoCard, winningNumber int) int {
	n := len(bingoCard[0])
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			number := bingoCard[i][j]
			if number == winningNumber {
				bingoCard[i][j] = -1
				if checkIfWinner(*bingoCard, i, j) {
					return sumOfUnmarkedNumbers(bingoCard)
				}
			}
		}
	}
	return -1
}

func readLines(path string) ([]BingoCard, []int) {
	file, _ := os.OpenFile(path, os.O_RDWR, 0600)
	defer file.Close()
	rawBytes, _ := ioutil.ReadAll(file)
	lines := strings.Split(string(rawBytes), "\n")
	var winningNumbers = Map(strings.Split(lines[0], ","), strconv.Atoi)
	var bingoCards []BingoCard
	for i := 1; i < len(lines); {
		if len(lines[i]) == 0 {
			i += 1
			continue
		}
		var bingoCard = constructBingoCard(lines, i)
		bingoCards = append(bingoCards, bingoCard)
		i += 5
	}
	return bingoCards, winningNumbers
}

func playBingo(bingoCards []BingoCard, winningNumbers []int) (int, int) {
	for _, winningNumber := range winningNumbers {
		for _, bingoCard := range bingoCards {
			result := checkCard(&bingoCard, winningNumber)
			if result != -1 {
				return result, winningNumber
			}
		}
	}
	return -1, -1
}

func indexOf(arr []int, value int) int {
	for k, v := range arr {
		if v == value {
			return k
		}
	}
	return -1
}

func findLastWinningCard(bingoCards []BingoCard, winningNumbers []int) int {
	lastIndex := -1
	lastResult := -1
	for _, bingoCard := range bingoCards {
		sumOfUnmarked, winningNumber := playBingo([]BingoCard{bingoCard}, winningNumbers)
		indexOfWinningNumber := indexOf(winningNumbers, winningNumber)
		if indexOf(winningNumbers, winningNumber) > lastIndex {
			lastResult = sumOfUnmarked * winningNumber
			lastIndex = indexOfWinningNumber
		}
	}
	return lastResult
}

func main() {
	bingoCards, winningNumbers := readLines("day_4/day_4.txt")
	sumOfUnmarked, winningNumber := playBingo(bingoCards, winningNumbers)
	println(sumOfUnmarked * winningNumber)
	bingoCards, winningNumbers = readLines("day_4/day_4.txt")
	println(findLastWinningCard(bingoCards, winningNumbers))

}
