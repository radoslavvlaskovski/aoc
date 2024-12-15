package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func solveTen() int {
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
        sum += getHikes(grid, x)
    }

    return sum
}

func getHikes(grid [][]int, start []int) int {
    count := 0
    n := len(grid)
    m := len(grid[0])
    visited := [][]bool{}

    for range n {
       rowVisited := []bool{}
       for range m {
           rowVisited = append(rowVisited, false)
       }
       visited = append(visited, rowVisited)
    }

    stack := make(chan []int, n * m)

    stack <- start

    for len(stack) != 0 {
       curr := <-stack
       x := curr[0]
       y := curr[1]

       if visited[x][y] {
           continue
       }

       fmt.Println(curr)

       visited[x][y] = true

       if grid[x][y] == 9 {
           count++
	   continue
       }

       if isInBounds([]int{x - 1, y}, n, m) && grid[x - 1][y] == grid[x][y] + 1 {
            stack <- []int{x - 1, y}
       }

       if isInBounds([]int{x + 1, y}, n, m) && grid[x + 1][y] == grid[x][y] + 1 {
            stack <- []int{x + 1, y}
       }

       if isInBounds([]int{x, y - 1}, n, m) && grid[x][y - 1] == grid[x][y] + 1 {
            stack <- []int{x, y - 1}
       }

       if isInBounds([]int{x, y + 1}, n, m) && grid[x][y + 1] == grid[x][y] + 1 {
            stack <- []int{x, y + 1}
       }

    }

    fmt.Println("start: ", start, " c: ", count)

    return count

}
