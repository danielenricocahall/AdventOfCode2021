package main

import (
	. "AdventOfCode/utils"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Foo struct {
	updated bool
}

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

func applyPolymerRules(polymerTemplate string, rules map[string]string, steps int) map[string]int {
	counter := make(map[string]int)

	for i := 0; i <= len(polymerTemplate)-2; i++ {
		pair := polymerTemplate[i : i+2]
		counter[pair] += 1
	}
	fmt.Println(counter)

	for step := 0; step < steps; step++ {
		println("STEP: " + strconv.Itoa(step))
		newCounter := make(map[string]int)
		updated := make(map[string]bool)
		for pair, count := range counter {
			if !updated[rules[pair]+string(pair[1])] {
				newCounter[rules[pair]+string(pair[1])] = count
				updated[rules[pair]+string(pair[1])] = true
			} else {
				newCounter[rules[pair]+string(pair[1])] += count
			}
			if !updated[string(pair[0])+rules[pair]] {
				newCounter[string(pair[0])+rules[pair]] = count
				updated[string(pair[0])+rules[pair]] = true

			} else {
				newCounter[string(pair[0])+rules[pair]] += count
			}
		}
		counter = newCounter
		fmt.Println(counter)
	}
	return counter
}

func countOccurrences(pairCounter map[string]int) {
	characterCounter := make(map[string]int)
	length := 0
	for pair, count := range pairCounter {
		characterCounter[string(pair[0])] += count
		length += count
	}
	fmt.Println(characterCounter)
	v := make([]int, 0, len(characterCounter))

	for _, value := range characterCounter {
		v = append(v, value)
	}

	println(ComputeMax(v) - ComputeMin(v) - 1)

}

func main() {
	rules, polymerTemplate := readLines("day_14/data.txt")
	counter := applyPolymerRules(polymerTemplate, rules, 40)
	countOccurrences(counter)

}
