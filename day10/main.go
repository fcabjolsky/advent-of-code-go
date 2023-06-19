package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NOOP = "noop"
	ADDX = "addx"
)

type Cpu struct {
	Cycle                 int
	X                     int
	currentOperationValue int
	currentOperationCycle int
	singalStrengths       map[int]int
	op                    string
}

func NewCpu() *Cpu {
	return &Cpu{
		Cycle:                 0,
		X:                     1,
		currentOperationValue: 0,
		currentOperationCycle: 0,
		singalStrengths:       map[int]int{},
	}
}

func (cpu *Cpu) StoreSignalStrength() {
	if cpu.Cycle == 20 || (cpu.Cycle-20)%40 == 0 {
		cpu.singalStrengths[cpu.Cycle] = cpu.Cycle * cpu.X
	}

}

func (cpu *Cpu) Next() {
	cpu.currentOperationCycle++
    cpu.Cycle++
	cpu.StoreSignalStrength()
	if cpu.currentOperationCycle == 2 {
		cpu.currentOperationCycle = 0
		cpu.X += cpu.currentOperationValue
		cpu.op = ""
	}
	splitted := strings.Split(cpu.op, " ")
	command := splitted[0]
	if command == ADDX {
		v, _ := strconv.Atoi(splitted[1])
		cpu.currentOperationValue = v
		cpu.Next()
	} else { 
        cpu.currentOperationCycle = 0
    }

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		file = os.Stdin
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    result := run(scanner)
    fmt.Println("Part 1:", result)

}

func run(scanner *bufio.Scanner) int {
    cpu := NewCpu()
	for scanner.Scan() {
		line := scanner.Text()
        cpu.op = line
        cpu.Next()
	}
    sum := 0
    for _, v := range cpu.singalStrengths {
        sum += v
    }
    return sum
}
