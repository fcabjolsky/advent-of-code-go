package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		file = os.Stdin
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	inputMap := processInput(scanner)
	part1 := part1(inputMap)
	fmt.Println(part1)

}

func processInput(scanner *bufio.Scanner) [][]int {
	inputMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		currentRow := []int{}
		for _, v := range line {
			tree, _ := strconv.Atoi(string(v))
			currentRow = append(currentRow, tree)
		}
		inputMap = append(inputMap, currentRow)
	}
	return inputMap
}

func checkRow(inputMap *[][]int, start, end, col, curr int) bool {
	for i := start; i < end; i++ {
		if (*inputMap)[col][i] >= curr {
			return false
		}
	}
	return true
}
func checkCol(inputMap *[][]int, start, end, row, curr int) bool {
	for i := start; i < end; i++ {
		if (*inputMap)[i][row] >= curr {
			return false
		}
	}
	return true
}

func part1(inputMap [][]int) int {
	total := ((len(inputMap) * 2) + (len(inputMap[0]) * 2)) - 4
	for i := 1; i < len(inputMap)-1; i++ {
		for j := 1; j < len(inputMap[i])-1; j++ {
			visibleT := checkCol(&inputMap, 0, i, j, inputMap[i][j]) 
			visibleB := checkCol(&inputMap, i + 1, len(inputMap), j, inputMap[i][j]) 
			visibleR := checkRow(&inputMap, j + 1, len(inputMap), i, inputMap[i][j]) 
			visibleL := checkRow(&inputMap, 0, j, i, inputMap[i][j]) 
			if visibleT || visibleB || visibleL || visibleR {
				total++
			}
		}

	}
	return total
}
