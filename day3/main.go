package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getMatch(firstComp string, secondComp string) rune {
	for _, letter := range firstComp {
		if strings.ContainsRune(secondComp, letter) {
			return letter
		}
	}
	return ' '
}

func getPointsForElement(element rune) int {
	if int(element) > 96 {
		return int(element) - 96
	}
	return (int(element) - 64) + 26
}
func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		halfLineLen := len(line) / 2
		firstHalf := line[:halfLineLen]
		secondHalf := line[halfLineLen:]
		element := getMatch(firstHalf, secondHalf)
		score += getPointsForElement(element)
	}
	fmt.Println(score)
}

func main() {
    Part2()
}
