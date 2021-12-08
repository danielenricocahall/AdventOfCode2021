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

func simulateLanternfishGrowth(
	lanternFishStartingStates []int,
	days int) map[int]int {
	lanternFishStates := lanternFishStartingStates
	foo := make(map[int]int)
	for i := 0; i <= 8; i++ {
		foo[i] = 0
	}
	for _, state := range lanternFishStates {
		foo[state] += 1
	}
	keys := make([]int, 0, len(foo))
	for k := range foo {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for day := 0; day < days; day++ {
		bar := foo[0]
		for _, lanternFishState := range keys {
			count := foo[lanternFishState]
			if lanternFishState > 0 {
				foo[lanternFishState-1] = count
			}
			foo[lanternFishState] = 0
		}
		foo[6] += bar
		foo[8] = bar
	}
	for k, _ := range keys {
		fmt.Println("Key:", k, "=>", "Element:", foo[k])
	}
	return foo
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
