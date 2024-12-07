package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func solveSix() int {
	file, err := os.Open("input6")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    count := 0
    scanner := bufio.NewScanner(file)

    block := byte('#')
    layout := []string{}
    visited := [][]bool{}

    for scanner.Scan() {
        line := scanner.Text()
        layout = append(layout, line)

	rowVisited := []bool{}
	for _, i := range line {
		i = i
	     rowVisited = append(rowVisited, false)
	}
	visited = append(visited, rowVisited)
    }

    n := len(layout)
    m := len(layout[0])
    curr := findStart(layout)
    dir := []int{-1, 0}

    fmt.Println("start: ", curr)

    for isInBounds(curr, n, m) {
	    fmt.Println(curr)
	    if !visited[curr[0]][curr[1]] {
	        count++
	    }
	    visited[curr[0]][curr[1]] = true
	    next := []int{curr[0] + dir[0], curr[1] + dir[1]}

	    if isInBounds(next, n, m) && layout[next[0]][next[1]] == block {
	        // rotate
                dir = rotate(dir)

		continue
	    }
	    
	    curr = next
    }

    return count;
}

func rotate(dir []int) []int {
	newDir := []int{dir[1], dir[0]}
	if dir[0] != 0 {
	    newDir[1] = dir[0] * -1
	}
	
	return newDir
}

func isInBounds(curr []int, n int, m int) bool {
	return curr[0] >= 0 && curr[1] >= 0 && curr[0] < n && curr[1] < m
}

func findStart(layout []string) []int {
    start := byte('^')
    for i, row := range layout {
	    for j, _ := range row {
	    if row[j] == start {
	       return []int{i, j}
	    }
	}
    }

    return nil
}

