package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func solveTwelve() int {
    file, err := os.Open("test")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var grid []string
    var visited [][]bool

    for scanner.Scan() {
         line := scanner.Text()

	 grid = append(grid, line)
	 var rowV []bool
	 for range len(line) {
	     rowV = append(rowV, false)
	 }

	 visited = append(visited, rowV)
    }

    fmt.Println(grid)

    n := len(grid)
    m := len(grid[0])

    sum := 0

    for i := range n {
        for j := range m {
	    if !visited[i][j] {
	       sum += visitGarden(grid, []int{i, j}, &visited)
	    }
	}
    }

    return sum
}


func visitGarden(grid []string, start []int, visited *[][]bool) int {
    a := 0
    p := 0

    c := grid[start[0]][start[1]]

    stack := [][]int{start}

    for len(stack) != 0 {
	curr := stack[0]
        x := curr[0]
	y := curr[1]

	stack = stack[1:]

	if (*visited)[x][y] {
	   continue
	}

	(*visited)[x][y] = true
	a++

	if matchesL(grid, c, x - 1, y) {
	    stack = append(stack, []int{x - 1, y})
	} else {
	    if !(matchesL(grid, c, x, y - 1) && !(*visited)[x][y - 1] && !matchesL(grid, c, x - 1, y - 1)) && !(matchesL(grid, c, x, y + 1) && !(*visited)[x][y + 1] && !matchesL(grid, c, x - 1, y + 1))  {
	       p++
	    }
	}

	if matchesL(grid, c, x + 1, y) {
            stack = append(stack, []int{x + 1, y})
        } else {
           if !(matchesL(grid, c, x, y - 1) && !(*visited)[x][y - 1] && !matchesL(grid, c, x + 1, y - 1)) && !(matchesL(grid, c, x, y + 1) && !(*visited)[x][y + 1] && !matchesL(grid, c, x + 1, y + 1))  {
               p++
            }
	}

	if matchesL(grid, c, x, y - 1) {
            stack = append(stack, []int{x, y - 1})
        } else {
             if !(matchesL(grid, c, x - 1, y) && !(*visited)[x - 1][y] && !matchesL(grid, c, x - 1, y - 1)) && !(matchesL(grid, c, x + 1, y) && !(*visited)[x + 1][y] && !matchesL(grid, c, x + 1, y - 1))  {
               p++
            }
	}

	if matchesL(grid, c, x, y + 1) {
            stack = append(stack, []int{x, y + 1})
        } else {
            if !(matchesL(grid, c, x - 1, y) && !(*visited)[x - 1][y] && !matchesL(grid, c, x - 1, y + 1)) && !(matchesL(grid, c, x + 1, y) && !(*visited)[x + 1][y] && !matchesL(grid, c, x + 1, y + 1))  {
               p++
            }
	}
    }

    fmt.Println("start: ", start, "char: ", string(c), " a: ", a, " p:", p)

    return a * p
}

func matchesL(grid []string, c byte, x int, y int) bool {
    n := len(grid)
    m := len(grid[0])

    return isInBounds([]int{x, y}, n, m) && grid[x][y] == c
}
