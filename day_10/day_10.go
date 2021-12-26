package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}

}
func readLines(path string) []string {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var results []string
	for scanner.Scan() {
		text := scanner.Text()
		results = append(results, text)
	}
	return results
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func parseNavigationSubsystem(navigationSystem []string) ([]string, []string) {
	symbols := []string{"{", "(", "[", "<"}
	op := map[string]string{
		symbols[0]: "}",
		symbols[1]: ")",
		symbols[2]: "]",
		symbols[3]: ">"}
	var syntaxErrorResults []string
	var incompleteResults []string
	for _, line := range navigationSystem {
		var stack Stack
		syntaxError := false
		for _, char := range strings.Split(line, "") {
			if stringInSlice(char, symbols) {
				stack.Push(char)
			} else {
				last, somethingFound := stack.Pop()
				if somethingFound {
					closing := op[last]
					if char != closing {
						println("Expected " + closing + " but found " + char + " instead")
						syntaxErrorResults = append(syntaxErrorResults, char)
						syntaxError = true
						break
					}
				}
			}
		}
		if syntaxError {
			continue
		}
		incompleteResult := ""
		for !stack.IsEmpty() {
			last, _ := stack.Pop()
			closing := op[last]
			incompleteResult += closing
		}
		incompleteResults = append(incompleteResults, incompleteResult)
	}
	return syntaxErrorResults, incompleteResults
}

func computeSyntaxScore(results []string) int {
	scoreTable := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	score := 0
	for _, result := range results {
		score += scoreTable[result]
	}
	return score
}

func computeIncompleteScore(results []string) []int {
	scoreTable := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var scores []int
	for _, result := range results {
		totalScore := 0
		for _, char := range strings.Split(result, "") {
			totalScore *= 5
			totalScore += scoreTable[char]
		}
		scores = append(scores, totalScore)
	}
	return scores
}
func main() {
	navigationSubsystem := readLines("day_10/data.txt")
	syntaxErrorResults, incompleteResults := parseNavigationSubsystem(navigationSubsystem)
	println(computeSyntaxScore(syntaxErrorResults))
	incompleteScores := computeIncompleteScore(incompleteResults)
	sort.Ints(incompleteScores)
	println(incompleteScores[len(incompleteScores)/2])

}
