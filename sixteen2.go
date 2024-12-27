package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func solveSixteen2() int {
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

    visited := [][]int{}
    for range len(grid) {
	row := []int{}
        for range len(grid[0]) {
	    row = append(row, 0)
	}
	visited = append(visited, row)
    }
    visited[s[0]][s[1]] = 1

    fmt.Println(s, e)

    path := [][]int{s}
    dfs16(grid, s, e, []int{0, 1}, 0, &path, 79404, &visited)

    countVisited := 0

    for i := range len(visited) {
	    for j := range len(visited[0]) {
	    if visited[i][j] > 0 {
	        countVisited++
	    }
	}
    }

    fmt.Println(visited)

    return countVisited
}

func dfs16(grid [][]int, curr []int, e []int, dir []int, score int, path *[][]int, t int, visited *[][]int) {
    if score > t {
       return
    }
    if curr[0] == e[0] && curr[1] == e[1] {
	    for i := range len(*path) {
	        x := (*path)[i]
		(*visited)[x[0]][x[1]] = 1
	    }
	return
    }

    u := checkInDir2(grid, curr, dir)
    if u {
        nextX := curr[0] + dir[0]
        nextY := curr[1] + dir[1]
        next := []int{nextX, nextY}
        nextDir := []int{dir[0], dir[1]}
	*path = append(*path, next)
	dfs16(grid, next, e, nextDir, score + 1, path, t, visited)
	*path = (*path)[:len(*path)-1]
    }

    dir = rotate(dir)
    u = checkInDir2(grid, curr, dir)
    if u {
        nextX := curr[0] + dir[0]
        nextY := curr[1] + dir[1]
        next := []int{nextX, nextY}
        nextDir := []int{dir[0], dir[1]}
        *path = append(*path, next)
        dfs16(grid, next, e, nextDir, score + 1001, path, t, visited)
	*path = (*path)[:len(*path)-1]
    }

    dir = rotate(dir)
    u = checkInDir2(grid, curr, dir)
    if u {
        nextX := curr[0] + dir[0]
        nextY := curr[1] + dir[1]
        next := []int{nextX, nextY}
        nextDir := []int{dir[0], dir[1]}
        *path = append(*path, next)
        dfs16(grid, next, e, nextDir, score + 2001, path, t, visited)
	*path = (*path)[:len(*path)-1]
    }

    dir = rotate(dir)
    u = checkInDir2(grid, curr, dir)
    if u {
        nextX := curr[0] + dir[0]
        nextY := curr[1] + dir[1]
        next := []int{nextX, nextY}
        nextDir := []int{dir[0], dir[1]}
        *path = append(*path, next)
        dfs16(grid, next, e, nextDir, score + 1001, path, t, visited)
	*path = (*path)[:len(*path)-1]
    }
}

func checkInDir2(grid [][]int, curr []int, dir []int) bool {
    nextX := curr[0] + dir[0]
    nextY := curr[1] + dir[1]

    return grid[nextX][nextY] != 1
}

