package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
    "sort"
)

func solveFifteen2() int {
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

    grid, moves, curr := parseInputSeventeen2(input)

    for _, m := range moves {
	    //print17(grid, curr)
	    fmt.Println("moving: ", m)
        x := curr[0] + m[0]
	y := curr[1] + m[1]

	if grid[x][y] == 1 {
	    continue
	}

	if grid[x][y] == 2 || grid[x][y] == 3 {

            if m[0] != 0 {
	        aX := x
	        var s, e int
	        if grid[x][y] == 2 {
	           s = y
	           e = y + 1
	        } else {
	           s = y - 1
	           e = y
	        }

		shouldCopy := false
		zones := [][]int{}
		zones = append(zones, []int{s, e})

		for true {
		    aX += m[0]
		    blockedRow := false

		    lastZone := zones[len(zones) - 1]
		    nextZone := []int{}
		    for _, i := range lastZone  {
		        if grid[aX][i] == 1 {
			   blockedRow = true
			   break
			}
			if grid[aX][i] == 2 {
			   nextZone = append(nextZone, i, i + 1)
			}
			if grid[aX][i] == 3 {
			    nextZone = append(nextZone, i - 1, i)
			}
		    }

		    if blockedRow {
		        break
		    }

		    if len(nextZone) == 0 {
		        shouldCopy = true
			break
		    }

		    zones = append(zones, nextZone)
		}

		if shouldCopy {
		   i := aX
		   for len(zones) > 0 {
			  
			r := zones[len(zones) - 1]
			sort.Ints(r)
			for rI, j := range r {
			    if rI > 0 && r[rI] == r[rI - 1] {
				    continue
			    }
			    grid[i][j] = grid[i-m[0]][j]
			    grid[i-m[0]][j] = 0
			}

			zones = zones[:len(zones)-1]
			i -= m[0]
		   }
		} else {
		   continue
		}

            } else {
	      b := y

	      for grid[x][b] == 2 || grid[x][b] == 3 {
		  b += m[1]
	      }

	      if grid[x][b] == 0 {
		 for b != y {
	             grid[x][b] = grid[x][b - m[1]]
		     b -= m[1]
		 }
	      } else {
	         continue
	      }
	   }
	}

	grid[x][y] = 0
	curr = []int{x, y}
    }

    print17(grid, curr)

    return sum17(grid)
}

func parseInputSeventeen2(lines []string) ([][]int, [][]int, []int) {
    input := [][]int{}
    var start []int
    moves := [][]int{}

    for _, line := range lines {
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
        for _, c := range line {
	    if byte(c) == byte('#') {
	        row = append(row, 1, 1)
	    } else if byte(c) == byte('O') {
	        row = append(row, 2, 3)
	    } else {
	        row = append(row, 0, 0)
	    }
	    if byte(c) == byte('@') {
	        start = []int{len(input), len(row) - 2}
	    }
	}
	input = append(input, row)
    }

    return input, moves, start
}

func print17(grid [][]int, curr []int) {
    for i := range len(grid) {

       for j, num := range grid[i] {
	 if i == curr[0] && j == curr[1] {
	    fmt.Print(" @ ")
	    continue
	 }
         if num == 0 {
             fmt.Print(" . ")
         } else if num == 1 {
             fmt.Print(" # ")
         } else if num == 2 {
 	    fmt.Print(" [ ")
	} else {
	    fmt.Print(" ] ")
	}
       }
       fmt.Println()
    }
}

