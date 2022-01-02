package main

import (
	. "AdventOfCode/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name  string
	edges []*Node
}

func findNodeInGraph(graph *[]*Node, nodeName string) *Node {
	for _, node := range *graph {
		if node.name == nodeName {
			return node
		}
	}
	return nil
}

func nodeInGraph(graph *[]*Node, nodeName string) bool {
	return findNodeInGraph(graph, nodeName) != nil
}

func addOrMergeNodes(graph *[]*Node, newNodeBegin *Node, newNodeEnd *Node) {
	if nodeInGraph(graph, newNodeBegin.name) {
		newNodeBegin = findNodeInGraph(graph, newNodeBegin.name)
	} else {
		*graph = append(*graph, newNodeBegin)
	}
	if nodeInGraph(graph, newNodeEnd.name) {
		newNodeEnd = findNodeInGraph(graph, newNodeEnd.name)
	} else {
		*graph = append(*graph, newNodeEnd)
	}
	newNodeBegin.edges = append(newNodeBegin.edges, newNodeEnd)
	newNodeEnd.edges = append(newNodeEnd.edges, newNodeBegin)
}

func readLines(path string) *[]*Node {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	graph := make([]*Node, 0)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, "-")
		begin := Node{name: splitted[0]}
		end := Node{name: splitted[1]}
		addOrMergeNodes(&graph, &begin, &end)
	}
	return &graph
}

func printGraph(graph *[]*Node) {
	for i, node := range *graph {
		print(node.name)
		if i < len(*graph)-1 {
			print("->")
		}
	}
	println()
}

func findAllPaths(graph *[]*Node, maxVisitCountForSmallCaves int) {
	start := findNodeInGraph(graph, "start")
	var currentPath []*Node
	pathCount := 0
	counter := make(map[string]int)
	traverse(start, currentPath, &pathCount, counter, maxVisitCountForSmallCaves)
	println("Total Paths: " + strconv.Itoa(pathCount))
}

func smallCavesNotVisited(
	smallCaveVisitedCounter map[string]int,
	maxVisitCount int) bool {
	for node, _ := range smallCaveVisitedCounter {
		if !smallCaveNotVisited(smallCaveVisitedCounter, node, maxVisitCount) {
			return false
		}
	}
	return true
}

func getSmallCaveNames(smallCaveVisitCounter map[string]int) []string {
	var smallCaves []string
	for cave, _ := range smallCaveVisitCounter {
		smallCaves = append(smallCaves, cave)
	}
	return smallCaves
}

func isSmallCave(cave string, smallCaves []string) bool {
	for _, smallCave := range smallCaves {
		if cave == smallCave {
			return true
		}
	}
	return false
}

func smallCaveNotVisited(smallCaveVisitedCounter map[string]int,
	smallCave string,
	maxVisitCount int) bool {
	return smallCaveVisitedCounter[smallCave] < maxVisitCount
}

func traverse(cave *Node,
	currentPath []*Node,
	pathCount *int,
	smallCaveVisitedCounter map[string]int,
	maxVisitCount int) {

	currentPath = append(currentPath, cave)
	if (*cave).name == "end" {
		printGraph(&currentPath)
		*pathCount += 1
		return
	}
	if IsLowerCase((*cave).name) {
		smallCaveVisitedCounter[(*cave).name] += 1
	}
	smallCaves := getSmallCaveNames(smallCaveVisitedCounter)
	for _, nextCave := range (*cave).edges {
		if !isSmallCave((*nextCave).name, smallCaves) ||
			(isSmallCave((*nextCave).name, smallCaves) &&
				smallCavesNotVisited(smallCaveVisitedCounter, maxVisitCount) &&
				(*nextCave).name != "start") {
			traverse(nextCave, currentPath, pathCount, CopyMap(smallCaveVisitedCounter), maxVisitCount)
		}
	}
}

func main() {
	graph := readLines("day_12/data.txt")
	findAllPaths(graph, 2)
}
