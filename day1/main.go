package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main () {
    scanner := bufio.NewScanner(os.Stdin)
    elfsCalories := []int{0}
    i := 0
    for scanner.Scan() {
        text := scanner.Text()
        if strings.TrimSpace(text) == "" {
            elfsCalories = append(elfsCalories, 0)
            i++
            continue
        }
        calories, err := strconv.Atoi(text)
        if err != nil {
            fmt.Println("error converting to int")
        }

        elfsCalories[i] += calories
    }
    sort.Ints(elfsCalories)
    
    topn := 3
    sum := 0
    for _, value := range elfsCalories[len(elfsCalories) - topn:] {
        sum += value
    }


    fmt.Println("max::", sum)
}
