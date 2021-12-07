package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Map(vs []string, f func(string) (int, error)) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i], _ = f(v)
	}
	return vsm
}

func readLines(path string) []int {
	file, _ := os.OpenFile(path, os.O_RDWR, 0600)
	defer file.Close()
	rawBytes, _ := ioutil.ReadAll(file)
	lines := strings.Split(string(rawBytes), "\n")
	lanternfishStartingStates := Map(strings.Split(lines[0], ","), strconv.Atoi)
	return lanternfishStartingStates
}

func simulateLanternfishGrowth(
	lanternFishStartingStates []int,
	days int) []int {
	lanternFishStates := lanternFishStartingStates
	for day := 0; day < days; day++ {
		for i, lanternFishState := range lanternFishStates {
			if lanternFishState == 0 {
				lanternFishStates[i] = 6
				lanternFishStates = append(lanternFishStates, 8)
			} else {
				lanternFishStates[i] = lanternFishState - 1
			}
		}
	}
	return lanternFishStates
}

func main() {
	lanternFishStartingStates := readLines("day_6/day_6_test.txt")
	results := simulateLanternfishGrowth(lanternFishStartingStates, 256)
	println(len(results))
}
