package main

import (
	"bufio"
	"os"
	"strings"
)

func readLines(path string) []string {

	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var codes []string
	for scanner.Scan() {
		text := scanner.Text()
		code := strings.Split(text, " | ")[1]
		codes = append(codes, code)
	}
	return codes
}

func countOutputValueOccurrences(codes []string) int {
	occurrences := 0
	for _, code := range codes {
		for _, x := range strings.Split(code, " ") {
			if (len(x) >= 2 && len(x) <= 4) || len(x) == 7 {
				occurrences += 1
			}
		}
	}
	return occurrences
}

func main() {
	codes := readLines("day_8/data.txt")
	println(countOutputValueOccurrences(codes))
}
