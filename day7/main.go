package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		file = os.Stdin
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	part1 := part1(scanner)
	fmt.Println(part1)

}

type Node struct {
	Parent *Node
	Childs []*Node
	Path   string
	Size   int
}

func NewNode(path string) *Node {
	return &Node{
		Childs: make([]*Node, 0),
		Path:   path,
		Size:   0,
	}
}

func (n *Node) setParent(parent *Node) {
	n.Parent = parent
}

// func (n *Node) addChild(child *Node) {
// 	n.Childs = append(n.Childs, child)
// }

func (n *Node) addSize(size int) {
	n.Size += size
}

func processCommand(currentNode *Node, command string, nodes *[]*Node) *Node {
	if strings.Contains(command, "cd") {
		dir := strings.Split(command, " ")[2]
		if !strings.Contains(dir, "..") {
			if dir != "/" {
				dir = currentNode.Path + dir
			}
			newNode := NewNode(dir)
			newNode.setParent(currentNode)
			currentNode = newNode
			*nodes = append(*nodes, currentNode)
		} else {
			currentNode.Parent.addSize(currentNode.Size)
			currentNode = currentNode.Parent
		}
	}
	return currentNode
}

func part1(scanner *bufio.Scanner) int {
	nodes := []*Node{}
	var currentNode *Node
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsRune(line, '$') {
			currentNode = processCommand(currentNode, line, &nodes)
		} else {
			if !strings.Contains(line, "dir") {
				size, _ := strconv.Atoi(strings.Split(line, " ")[0])
				currentNode.addSize(size)
			}
		}

	}

	for currentNode.Path != "/" {
		currentNode.Parent.addSize(currentNode.Size)
		currentNode = currentNode.Parent
	}

	var sum int = 0
	for _, v := range nodes {
		if v.Size <= 100000 {
			sum += v.Size
		}
	}
	return sum
}
