package main

import (
	. "AdventOfCode/utils"
	"bufio"
	"fmt"
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
	for _, node := range *graph {
		println(node.name)
		for _, edge := range node.edges {
			fmt.Println(*edge)
		}
	}
}

func findAllPaths(graph *[]*Node) {
	start := findNodeInGraph(graph, "start")
	previous := []*Node{}
	currentPath := []*Node{}
	pathCount := 0
	traverse(start, previous, currentPath, &pathCount)
	println("Total Paths: " + strconv.Itoa(pathCount))
}

func traverse(node *Node, subGraph []*Node, currentPath []*Node, pathCount *int) {
	if (*node).name == "end" {
		for _, foo := range currentPath {
			print(foo.name + "->")
		}
		print((*node).name)
		println()
		*pathCount += 1
		return
	}
	if IsLowerCase((*node).name) || (*node).name == "start" {
		subGraph = append(subGraph, node)
	}
	currentPath = append(currentPath, node)
	for _, neighbor := range (*node).edges {
		if !nodeInGraph(&subGraph, (*neighbor).name) {
			traverse(neighbor, subGraph, currentPath, pathCount)
		}
	}
}

func main() {
	graph := readLines("day_12/data.txt")
	findAllPaths(graph)
}
