package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

func initializeCounter(lanternFishStartingStates []int) (map[int]int, []int) {
	stateCounter := make(map[int]int)
	for i := 0; i <= 8; i++ {
		stateCounter[i] = 0
	}
	for _, state := range lanternFishStartingStates {
		stateCounter[state] += 1
	}
	lanternFishStates := make([]int, 0, len(stateCounter))
	for k := range stateCounter {
		lanternFishStates = append(lanternFishStates, k)
	}
	sort.Ints(lanternFishStates)

	return stateCounter, lanternFishStates
}

func simulateLanternfishGrowth(
	lanternFishStartingStates []int,
	days int) map[int]int {
	counter, lanternFishStates := initializeCounter(lanternFishStartingStates)
	for day := 0; day < days; day++ {
		lanternFishAboutToSpawn := counter[0]
		for _, lanternFishState := range lanternFishStates {
			count := counter[lanternFishState]
			if lanternFishState > 0 {
				counter[lanternFishState-1] = count
			}
			counter[lanternFishState] = 0
		}
		counter[6] += lanternFishAboutToSpawn
		counter[8] = lanternFishAboutToSpawn
	}
	for k, _ := range lanternFishStates {
		fmt.Println("Lantern Fish State:", k, "=>", "Count:", counter[k])
	}
	return counter
}

func main() {
	lanternFishStartingStates := readLines("day_6/day_6.txt")
	results := simulateLanternfishGrowth(lanternFishStartingStates, 256)
	acc := 0
	for _, v := range results {
		acc += v
	}
	println(acc)
}
