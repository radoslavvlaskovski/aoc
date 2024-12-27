package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
)

func solveFifteen() int {
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

    grid, moves, curr := parseInputSeventeen(input)

    for _, m := range moves {
        x := curr[0] + m[0]
	y := curr[1] + m[1]

	if grid[x][y] == 1 {
	    continue
	}

	if grid[x][y] == 2 {
		a, b := x, y

	    for grid[a][b] == 2 {
	        a += m[0]
		b += m[1]
	    }

	    if grid[a][b] == 0 {
	        grid[a][b] = 2
		grid[x][y] = 0
	    } else {
	        continue
	    }
	}

	curr = []int{x, y}
    }

    return sum17(grid)
}

func sum17(grid [][]int) int {
    sum := 0
    for i := range len(grid) {
        for j := range len(grid[0]) {
	    if grid[i][j] == 2 {
	        sum += 100 * i + j 
	    }
	}
    }

    return sum
}

func parseInputSeventeen(lines []string) ([][]int, [][]int, []int) {
    input := [][]int{}
    var start []int
    moves := [][]int{}

    for i, line := range lines {
	    if len(line) == 0  {
	        continue
	    }
	if !strings.Contains(line, "#") {
		for _, c := range line {
		   if byte(c) == byte('>') {
		       moves = append(moves, []int{0, 1})
		   }
		   if byte(c) == byte('<') {
                       moves = append(moves, []int{0, -1})
                   }
		   if byte(c) == byte('^') {
                       moves = append(moves, []int{-1, 0})
                   }
		   if byte(c) == byte('v') {
                       moves = append(moves, []int{1, 0})
                   }
		}
            continue
	}

	row := []int{}
        for j, c := range line {
	    if byte(c) == byte('#') {
	        row = append(row, 1)
	    } else if byte(c) == byte('O') {
	        row = append(row, 2)
	    } else {
	        row = append(row, 0)
	    }
	    if byte(c) == byte('@') {
	        start = []int{i, j}
	    }
	}
	input = append(input, row)
    }

    return input, moves, start
}
