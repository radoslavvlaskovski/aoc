package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func solveSix2() int {
	file, err := os.Open("input6")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    layout := []string{}

    for scanner.Scan() {
	line := scanner.Text()
        layout = append(layout, line)
    }
    dir := []int{-1, 0}
    return solve(layout, dir, []int{-1, -1}, true)
}

func solve(layout []string, dir []int, start []int,  deepCheck bool) int {
    block := byte('#')
    count := 0
    visited := [][]int{}
    placementChecked := [][]bool{}

    for _, line := range layout {
        rowVisited := []int{}
	pC := []bool{}
        for range line {
             rowVisited = append(rowVisited, 0)
	     pC = append(pC, false)
        }
        visited = append(visited, rowVisited)
	placementChecked = append(placementChecked, pC)
    }

    n := len(layout)
    m := len(layout[0])
    if start[0] == -1 {
         start = findStart(layout)
    }
    curr := start

    for isInBounds(curr, n, m) {
	    if !deepCheck && visited[curr[0]][curr[1]] >= 5 {
	        return 1
	    }
            next := []int{curr[0] + dir[0], curr[1] + dir[1]}

            if isInBounds(next, n, m) && layout[next[0]][next[1]] == block {
                // rotate
                dir = rotate(dir)

                continue
            } else if isInBounds(next, n, m) && deepCheck && (start[0] != next[0] || start[1] != next[1]) && !placementChecked[next[0]][next[1]] && visited[next[0]][next[1]] == 0 {
		    potentialRot := rotate(dir)
		bS := []byte(layout[next[0]])
		bS[next[1]] = byte('#')
		layout[next[0]] = string(bS)

		sol := solve(layout, potentialRot, curr,  false)

		if sol == 1 {
			fmt.Println("adding at: ", next)
			placementChecked[next[0]][next[1]] = true
		}
		count += sol

		bS[next[1]] = byte('.')
		layout[next[0]] = string(bS)
            }
	    visited[curr[0]][curr[1]] += 1
            curr = next
    }

    return count;
}
