package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) (map[string]string, string) {
	rules := make(map[string]string)
	fileIO, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}
	polymerTemplate := ""
	lines := strings.Split(string(rawBytes), "\n")
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		if i == 0 {
			polymerTemplate = line
		} else {
			splitted := strings.Split(line, " -> ")
			rules[splitted[0]] = splitted[1]
		}
	}
	return rules, polymerTemplate
}

func applyPolymerRules(polymerTemplate string, rules map[string]string, steps int) string {
	newPolymer := polymerTemplate
	println(polymerTemplate)
	for step := 0; step < steps; step++ {
		print("STEP: " + strconv.Itoa(step))
		for i := 0; i <= len(polymerTemplate)-2; i++ {
			pair := polymerTemplate[i : i+2]
			newPolymer = newPolymer[:i+(i+1)] + rules[pair] + newPolymer[i+(i+1):]
		}
		println(newPolymer)
		polymerTemplate = newPolymer
	}
	return newPolymer
}

func countOccurrences(polymer string) {
	count := make(map[string]int)
	for _, char := range polymer {
		count[string(char)] += 1
	}
	fmt.Println(count)

}

func main() {
	rules, polymerTemplate := readLines("day_14/test.txt")
	polymer := applyPolymerRules(polymerTemplate, rules, 10)
	countOccurrences(polymer)

}
