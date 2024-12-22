package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
    "strconv"
)



func solveFourteen() int {
    file, err := os.Open("test")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    n := 101
    m := 103

    duration := 10403

    ps := [][]int{}
    vs := [][]int{}

    for scanner.Scan() {
         line := scanner.Text()

	 p, v := parsePosAndVel(line)
	 ps = append(ps, p)
	 vs = append(vs, v)
    }

    var grid [][]int
    var q1, q2, q3, q4 int

    for d := range duration {
	 fmt.Println("d: ", d)

	 grid = [][]int{}
         for range m {
              grid = append(grid, make([]int, n))
         }

	 for i := range len(ps) {
	     x := move(ps[i][0], vs[i][0], d, n)
             y := move(ps[i][1], vs[i][1], d, m)

	     grid[y][x] += 1
         }

	        q1 = countInArray(grid[0:m/2], 0, n/2)
		q2 = countInArray(grid[0:m/2], n/2 + 1, n)
		q3 = countInArray(grid[m/2+1:], 0, n/2)
		q4 = countInArray(grid[m/2+1:], n/2 + 1, n)

	        if treePattern(grid, m, n) {
		    fmt.Println("found")
		    for i := range len(grid) {
			 printArrayWithZerosAsDots(grid[i])
		    }
		    break
	       }
    }

    fmt.Println("q1: ", q1, "q2: ", q2, "q3: ", q3, "q4: ", q4)

    return q1 * q2 * q3 * q4
}

func treePattern(arr [][]int, m int, n int) bool {
    for x := range m {

	for y := range n {
		dX := x
		lY := y
		rY := y
		found := true

	for range 10 {
	    if (!isInBounds([]int{dX, lY}, m, n) || arr[dX][lY] == 0) {
		    found = false
		    break
	    }
            dX += 1
	    lY -= 1
	    rY += 1
           }

	   if found {
	       return true
	   }
       }
    }
    return false
}

func countInArray(arr [][]int, from int, to int) int {
    count := 0

    for i, _ := range arr {
        for j, _ := range arr[i] {
	    if j >= from && j < to {
	        count += arr[i][j]
	    }
	}
    }

    return count
}

func move(p int, v int, duration int, n int) int {
   x := (p + duration * v) % n

   if x != 0 && v < 0 && x < 0 {
      return n + x
   }

   return x
}

func parsePosAndVel(line string) ([]int, []int) {

	parts := strings.Split(line, " ")
	posStr := parts[0][2:]
	velStr := parts[1][2:]

	return parseIntPair(posStr), parseIntPair(velStr)
}


func parseIntPair(values string) []int {
    valuesStr := strings.Split(values, ",")
    left, _ := strconv.Atoi(valuesStr[0])
    right, _ := strconv.Atoi(valuesStr[1])

    return []int{left, right}
}

func printArrayWithZerosAsDots(arr []int) {
    for _, num := range arr {
        if num == 0 {
            fmt.Print(".")
        } else {
	    fmt.Print("X")
        }
    }
    fmt.Println() // New line after printing array
}

