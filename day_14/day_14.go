package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
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
	// NCNBCHB
	for step := 0; step < steps; step++ {
		println("Step: " + strconv.Itoa(step))
		var wg sync.WaitGroup
		wg.Add(len(polymerTemplate) - 1)
		changes := make([]string, len(polymerTemplate)-1)
		for i := 0; i <= len(polymerTemplate)-2; i++ {
			go func(i int) {
				defer wg.Done()
				pair := polymerTemplate[i : i+2]
				changes[i] = rules[pair]
			}(i)
		}
		wg.Wait()

		for i, change := range changes {
			newPolymer = newPolymer[:i+(i+1)] + change + newPolymer[i+(i+1):]
		}

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
	rules, polymerTemplate := readLines("day_14/data.txt")
	polymer := applyPolymerRules(polymerTemplate, rules, 40)
	countOccurrences(polymer)

}
