package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func solveSeven() int {
	file, err := os.Open("input7")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    sum := 0

    for scanner.Scan() {
	line := scanner.Text()
	result, values := parseEquation(line)

	fmt.Println(result, values)

	canBeSolved := solveEquation(values[0], result, values[1:])
	if canBeSolved {
		fmt.Println("solved")
	    sum += result
	}
    }

    return sum
}

func parseEquation(line string) (int, []int) {
	split := strings.Split(line, ":")
	result, _ := strconv.Atoi(split[0])
	
	strValues := strings.Split(split[1][1:], " ")
	values := []int{}
	
	for _, v := range strValues {
	    value, _ := strconv.Atoi(v)
	    values = append(values, value)
	}

	return result, values
}

func solveEquation(cur int, result int, remaining []int) bool {
    if cur > result {
        return false
    }

    if len(remaining) == 1 {
	  return cur + remaining[0] == result || cur * remaining[0] == result || concatNumbers(cur, remaining[0]) == result
    }

    return solveEquation(cur + remaining[0], result, remaining[1:]) || solveEquation(cur * remaining[0], result, remaining[1:]) || solveEquation(concatNumbers(cur, remaining[0]), result, remaining[1:])
}

func concatNumbers(left int, right int) int {
     leftStr := strconv.Itoa(left)
     rightStr := strconv.Itoa(right)
     v, _ := strconv.Atoi(leftStr + rightStr)
     return v
}
