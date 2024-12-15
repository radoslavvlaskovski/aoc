package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func solveTen2() int {
	file, err := os.Open("input10")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    sum := 0

    starts := [][]int{}
    grid := [][]int{}
    i := 0

    for scanner.Scan() {
         line := scanner.Text()

	 row := []int{}

	 for j, x := range line {
	      v, _ := byteToDigit(byte(x))
	      row = append(row, v)

	      if byte(x) != byte('.') && v == 0 {
	          starts = append(starts, []int{i, j})
	      }
	 }

	 grid = append(grid, row)
	 i++
    }
    fmt.Println(grid)

    for _, x := range starts {
        sum += getHikes2(grid, x)
    }

    return sum
}

func getHikes2(grid [][]int, start []int) int {
    count := 0
    n := len(grid)
    m := len(grid[0])
    counts := [][]int{}
    visited := [][]bool{}

    for i := range n {
       rowC := []int{}
       rowV := []bool{}
       for j := range m {
	   if grid[i][j] == 0 {
	       rowC = append(rowC, 1)
	   } else {
	       rowC = append(rowC, 0)
	   }
	   rowV = append(rowV, false)
       }
       counts = append(counts, rowC)
       visited = append(visited, rowV)
    }

    stack := [][]int{start}

    for len(stack) != 0 {
       curr := stack[0]
       stack = stack[1:]
       x := curr[0]
       y := curr[1]

       if visited[x][y] {
           continue
       }

       visited[x][y] = true

       if grid[x][y] == 9 {
           count += counts[x][y]
	   continue
       }

       if isInBounds([]int{x - 1, y}, n, m) && grid[x - 1][y] == grid[x][y] + 1 {
            stack = append(stack, []int{x - 1, y})
	    counts[x - 1][y] = counts[x - 1][y] + counts[x][y]
       }

       if isInBounds([]int{x + 1, y}, n, m) && grid[x + 1][y] == grid[x][y] + 1 {
            stack = append(stack, []int{x + 1, y})
	    counts[x + 1][y] = counts[x + 1][y] + counts[x][y]
       }

       if isInBounds([]int{x, y - 1}, n, m) && grid[x][y - 1] == grid[x][y] + 1 {
            stack = append(stack, []int{x, y - 1})
	    counts[x][y - 1] = counts[x][y - 1] + counts[x][y]
       }

       if isInBounds([]int{x, y + 1}, n, m) && grid[x][y + 1] == grid[x][y] + 1 {
            stack = append(stack, []int{x, y + 1})
	    counts[x][y + 1] = counts[x][y + 1] + counts[x][y]
       }

    }

    fmt.Println("start: ", start, " c: ", count)

    return count

}
