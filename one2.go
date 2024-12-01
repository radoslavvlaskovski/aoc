package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "strconv"
    "fmt"
)

func solveOne2() uint64 {
	file, err := os.Open("input1")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var valuesLeft []int
    valuesRight := make(map[int]uint64)

    for scanner.Scan() {
        line := scanner.Text()
	
	words := strings.Split(line, "   ")
        
	leftValue, _ := strconv.Atoi(words[0])
        rightValue, _ := strconv.Atoi(words[1])

	valuesLeft = append(valuesLeft, leftValue)

	value, exists := valuesRight[rightValue]
	if exists {
	    valuesRight[rightValue] = value + 1
	} else {
	    valuesRight[rightValue] = 1
	}
    }


    var sim uint64
    for _, leftV := range valuesLeft {
         value, exists := valuesRight[leftV]
	 if exists {
             fmt.Println(value * uint64(leftV))
	     sim += value * uint64(leftV)
	 }
    }

    return sim;
}
