package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var p = fmt.Println

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) PrintNode() {
	p("Node Value: ", node.Value)
	if node.Left != nil {
		node.PrintLeft()
	}
	if node.Right != nil {
		node.PrintRight()
	}
}

func (node *Node) PrintLeft() {
	p("Left value: ", node.Left.Value)
}

func (node *Node) PrintRight() {
	p("Right value: ", node.Right.Value)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := createNodesFromFile()
	for i := 0; i < len(data); i++ {
		data[i].PrintNode()
		p()
	}
}

func createNodesFromFile() []Node {
	f, err := os.Open("nodes.in")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines [][]string
	var N int
	var line []string
	for scanner.Scan() {
		if len(scanner.Text()) <= 1 {
			N, _ = strconv.Atoi(scanner.Text())
		} else {
			p(len(scanner.Text()))
			line = strings.Split(scanner.Text(), " ")
			lines = append(lines, line)
		}
	}
	p(lines)

	var nodes []Node = make([]Node, N)
	for i := 0; i < N; i++ {
		var Val, LeftIndex, RightIndex int
		Val, _ = strconv.Atoi(lines[i][0])
		LeftIndex, _ = strconv.Atoi(lines[i][1])
		RightIndex, _ = strconv.Atoi(lines[i][2])
		nodes[i].Value = Val
		if LeftIndex >= 0 {
			nodes[i].Left = &nodes[LeftIndex]
		}
		if RightIndex >= 0 {
			nodes[i].Right = &nodes[RightIndex]
		}
		i++

	}
	return nodes
}
