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
    elvesCalories := []int{0}
    i := 0
    for scanner.Scan() {
        text := scanner.Text()
        if strings.TrimSpace(text) == "" {
            elvesCalories = append(elvesCalories, 0)
            i++
            continue
        }
        calories, err := strconv.Atoi(text)
        if err != nil {
            fmt.Println("error converting to int")
        }

        elvesCalories[i] += calories
    }
    sort.Ints(elvesCalories)
    
    topn := 3
    sum := 0
    for _, value := range elvesCalories[len(elvesCalories) - topn:] {
        sum += value
    }


    fmt.Println("max::", sum)
}
