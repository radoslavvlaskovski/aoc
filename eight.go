package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "math"
)

func solveEight() int {
	file, err := os.Open("input8")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    count := 0
    isAntiNode := [][]bool{}
    cellLocs := make(map[byte][][]int)

    layout := []string{}
    
    for scanner.Scan() {
         line := scanner.Text()

	 layout = append(layout, line)

	 rowAnti := []bool{}

	 for range line {
            rowAnti = append(rowAnti, false)
	 }

	 isAntiNode = append(isAntiNode, rowAnti)
    }

    n := len(layout)
    m := len(layout[0])

    for i, line := range layout {
	for j, cellVRune := range line {
	    cellV := byte(cellVRune)
	    if cellV == byte('.') {
	        continue
	    }

	    prevCells, exists := cellLocs[cellV]

	    if !exists {
	        cellLocs[cellV] = [][]int{{i, j}}
		continue
	    }

	    for _, prevCell := range prevCells {
                nodes := createAntiNodeCand([]int{i, j}, prevCell, n, m)

		for _, node := range nodes {
		    isAntiNode[node[0]][node[1]] = true
		}
	    }

	    cellLocs[cellV] = append(cellLocs[cellV], []int{i, j})
	}
    }


    for i, _ := range isAntiNode {
        for j, _ := range isAntiNode[i] {
	    if isAntiNode[i][j] {
		    fmt.Println("Found at: ", i, j)
	        count++
	    }
	}
    }

    return count
}

func createAntiNodeCand(curr []int, prev []int, n int, m int) [][]int  {
    nodes := [][]int{}

    diffX := int(math.Abs(float64(curr[0] - prev[0])))
    diffY := int(math.Abs(float64(curr[1] - prev[1])))
    nodeOne := prev
    nodeTwo := curr

    if curr[1] < prev[1] {
	    for isInBounds(nodeOne, n, m) {
	        nodes = append(nodes, nodeOne)
		nodeOne = []int{nodeOne[0] - diffX, nodeOne[1] + diffY}
	    }

	    nodeTwo := curr

	    for isInBounds(nodeTwo, n, m) {
               nodes = append(nodes, nodeTwo)
	       nodeTwo = []int{nodeTwo[0] + diffX, nodeTwo[1] - diffY}
            }
    } else {
            for isInBounds(nodeOne, n, m) {
               nodes = append(nodes, nodeOne)
	       nodeOne = []int{nodeOne[0] - diffX, nodeOne[1] - diffY}
            }
            for isInBounds(nodeTwo, n, m) {
               nodes = append(nodes, nodeTwo)

	       nodeTwo = []int{nodeTwo[0] + diffX, nodeTwo[1] + diffY}
            }
    }

    return nodes
}

