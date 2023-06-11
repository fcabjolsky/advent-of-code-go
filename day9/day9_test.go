package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Day8(t *testing.T) {
	file, err := os.Open("input.txt")

	if err != nil {
		file = os.Stdin
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    result := part1(scanner)
    assert.Equal(t, 6181, result)
}

