package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "math"
)



func solveSixteen() int {
    file, err := os.Open("test")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    input := []string{}

    for scanner.Scan() {
         line := scanner.Text()
	 input = append(input, line)
    }

    grid, s, e := parseInputSixteen(input)

    fmt.Println(s, e)
    fmt.Println(grid)

    score := [][]int{}
    for i := range len(grid) {
	    row := []int{}
	    for range len(grid[i]) {
	       row = append(row, math.MaxInt64)
	    }
	    score = append(score, row)
    }
    score[s[0]][s[1]] = 0

    dir := []int{0, 1}
    stack := [][]int{s}
    stackD := [][]int{dir}

    for len(stack) != 0 {
	curr := stack[0]
	dir = stackD[0]
	stack = stack[1:]
	stackD = stackD[1:]

	if curr[0] == e[0] && curr[1] == e[1] {
	    continue
	}

	s, u := checkInDir(grid, score, curr, dir, 1)
	if u {
           nextX := curr[0] + dir[0]
           nextY := curr[1] + dir[1]
           score[nextX][nextY] = s
	   next := []int{nextX, nextY}
	   nextDir := []int{dir[0], dir[1]}
	   stack = append([][]int{next}, stack...)
	   stackD = append([][]int{nextDir}, stackD...)
	}

	rDir := rotate(dir)
	s, u = checkInDir(grid, score, curr, rDir, 1001)
        if u {
           nextX := curr[0] + rDir[0]
             nextY := curr[1] + rDir[1]
             score[nextX][nextY] = s
	     next := []int{nextX, nextY}
           nextDir := []int{rDir[0], rDir[1]}
           stack = append([][]int{next}, stack...)
           stackD = append([][]int{nextDir}, stackD...)
        }

	rDir = rotate(rDir)
	s, u = checkInDir(grid, score, curr, rDir, 2001)
	if u {
             nextX := curr[0] + rDir[0]
             nextY := curr[1] + rDir[1]
             score[nextX][nextY] = s
	     next := []int{nextX, nextY}
           nextDir := []int{rDir[0], rDir[1]}
           stack = append([][]int{next}, stack...)
           stackD = append([][]int{nextDir}, stackD...)
        }

        rDir = rotate(rDir)
        s, u = checkInDir(grid, score, curr, rDir, 1001)
        if u {
             nextX := curr[0] + rDir[0]
               nextY := curr[1] + rDir[1]
               score[nextX][nextY] = s
	       next := []int{nextX, nextY}
           nextDir := []int{rDir[0], rDir[1]}
           stack = append([][]int{next}, stack...)
           stackD = append([][]int{nextDir}, stackD...)
        }

    }

    for _, x := range score {
	    print16(x)
    }

    return score[e[0]][e[1]]
}

func checkInDir(grid [][]int, score [][]int, curr []int, dir []int, pen int) (int, bool) {
	currScore := score[curr[0]][curr[1]]
    nextX := curr[0] + dir[0]
    nextY := curr[1] + dir[1]

    if grid[nextX][nextY] == 1 {
	    return 0, false
    }

    return currScore + pen, currScore + pen < score[nextX][nextY]
}

func parseInputSixteen(lines []string) ([][]int, []int, []int) {
    input := [][]int{}
    var start []int
    var end []int

    for i, line := range lines {
	row := []int{}
        for j, c := range line {
	    if byte(c) == byte('#') {
	        row = append(row, 1)
	    } else {
	        row = append(row, 0)
	    }
	    if byte(c) == byte('S') {
	        start = []int{i, j}
	    }
	    if byte(c) == byte('E') {
	        end = []int{i, j}
	    }
	}
	input = append(input, row)
    }

    return input, start, end
}

func print16(arr []int) {
    for _, num := range arr {
        if num == math.MaxInt64 {
            fmt.Print(" . ")
        } else {
            fmt.Print(" ",num)
        }
    }
    fmt.Println() // New line after printing array
}
