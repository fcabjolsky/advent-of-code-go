package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Status int

const (
	Loss Status = iota
	Draw Status = iota + 2
	Win  Status = iota + 4
)

type GameOptions int

const (
	Rock GameOptions = iota + 1
	Paper
	Scissors
)

func getOponentOption(option string) GameOptions {
	getOponentOptionMap := map[string]GameOptions{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	return getOponentOptionMap[option]
}

func getMyOption(option string) GameOptions {
	myOptionMap := map[string]GameOptions{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
	return myOptionMap[option]
}

func getStatus(opSelection GameOptions, mySelection GameOptions) Status {
	switch opSelection {
	case Paper:
		{
			switch mySelection {
			case Rock:
				return Loss
			case Scissors:
				return Win
			}
		}
	case Rock:
		{
			switch mySelection {
			case Paper:
				return Win
			case Scissors:
				return Loss
			}
		}
	case Scissors:
		{
			switch mySelection {
			case Paper:
				return Loss
			case Rock:
				return Win
			}
		}
	}
	return Draw 
}

func day2() {
	scanner := bufio.NewScanner(os.Stdin)

	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")
		oponent := getOponentOption(splitted[0])
		mine := getMyOption(splitted[1])
		status := getStatus(oponent, mine)
		score += int(status) + int(mine)
	}
	fmt.Println(score)
}

func main() {
    // day2()
    Day2_2() 
}

