package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	// "os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint() *Point {
	return &Point{
		X: 0,
		Y: 0,
	}
}

func (p *Point) XModule(p2 *Point) float64 {
	return math.Pow(float64(p2.X-p.X), 2.0)
}
func (p *Point) YModule(p2 *Point) float64 {
	return math.Pow(float64(p2.Y-p.Y), 2.0)
}

func (p *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(p.XModule(p2) + p.YModule(p2))
}

func (p *Point) Move(direction string) {
	switch direction {
	case "U":
		p.X--
		return
	case "D":
		p.X++
	case "R":
		p.Y++
	case "L":
		p.Y--
	}
}

func (p *Point) MoveTo(direction string, p2 *Point) {
	switch direction {
	case "U":
		p.Y = p2.Y
		p.X = p2.X + 1
		return
	case "D":
		p.X = p2.X - 1
		p.Y = p2.Y
		return
	case "R":
		p.X = p2.X
		p.Y = p2.Y - 1
		return
	case "L":
		p.X = p2.X
		p.Y = p2.Y + 1
		return
	}
}
func (p *Point) Direction() string {
	return fmt.Sprintf("%d%d", p.X, p.Y)
}

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

func part1(scanner *bufio.Scanner) int {
	tail := NewPoint()
	head := NewPoint()
	visited := map[string]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, " ")
		direction := splitted[0]
		ammount, _ := strconv.Atoi(splitted[1])
		for i := 0; i < ammount; i++ {
			head.Move(direction)
			if head.Distance(tail) > 2 {
				tail.MoveTo(direction, head)
			} else if head.Distance(tail) == 2 {
				tail.Move(direction)
			}
			visited[tail.Direction()] = true
		}
	}
	return len(visited)
}
