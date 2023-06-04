package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getCommonElementForGroup(elfsRucksacks []string) rune {
	for _, element := range elfsRucksacks[0] {
		found := true
		for _, e := range elfsRucksacks[1:] {
			if !strings.ContainsRune(e, element) {
				found = false
				break
			}
		}
		if found {
			return element
		}
	}
	return ' '
}

func Part2() {
	scanner := bufio.NewScanner(os.Stdin)
	elfGroups := []string{}
	i := 0
	score := 0
    nGroup := 3
	for scanner.Scan() {
		line := scanner.Text()
		elfGroups = append(elfGroups, line)
		i++
		if i == nGroup {
			element := getCommonElementForGroup(elfGroups)
			score += getPointsForElement(element)
			elfGroups = []string{}
			i = 0
		}

	}
	fmt.Println(score)

}
