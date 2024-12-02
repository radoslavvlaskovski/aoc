package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "strconv"
    "fmt"
)

func solveTwo() int {
	file, err := os.Open("input2")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    count := 0
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
	
	words := strings.Split(line, " ")
        
       intList := make([]int, 0, len(words))

       for _, str := range words {
           num, _ := strconv.Atoi(str)
           intList = append(intList, num)
       }

       fmt.Println(intList)

       if safetyCheck(intList) {
           count++
       }
    }

    return count;
}


func safetyCheck(data []int) bool {
    if len(data) == 1 {
        return true
    }
    inc := data[1] > data[0]
    var left int
    var right int

    for i := 0; i < len(data); i++ {
        if i == 0 {
	    continue
	}
	if inc {
	    left = data[i]
	    right = data[i - 1]
	} else {
            left = data[i - 1]
	    right = data[i]
	}

	if !inLim(left, right) {
	     return false
	}
    }

    return true
}

func inLim(x, y int) bool {
    return x - y >= 1 && x - y <= 3
}

