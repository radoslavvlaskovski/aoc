package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "sort"
    "math"
    "strconv"
    "fmt"
)

func solveOne() uint64 {
	file, err := os.Open("input1")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var valuesLeft []int
    var valuesRight []int

    for scanner.Scan() {
        line := scanner.Text()
	
	words := strings.Split(line, "   ")
        
	leftValue, _ := strconv.Atoi(words[0])
        rightValue, _ := strconv.Atoi(words[1])

	valuesLeft = append(valuesLeft, leftValue)
	valuesRight = append(valuesRight, rightValue)
    }

    sort.Ints(valuesLeft)
    sort.Ints(valuesRight)


    var diff uint64
    for i, _ := range valuesLeft {
	fmt.Println(float64(valuesLeft[i] - valuesRight[i]))
        diff += uint64(math.Round(math.Abs(float64(valuesLeft[i] - valuesRight[i]))))
    }

    return diff;
}
