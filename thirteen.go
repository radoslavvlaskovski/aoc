package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func solveThirteen() int {
    file, err := os.Open("test")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    i := 0
    aX := 0
    aY := 0
    bX := 0
    bY := 0
    sum := 0

    for scanner.Scan() {
         line := scanner.Text()

	 if strings.Contains(line, "Button A") {
	    aX, aY = parseButtonValues(line)
	 }
         if strings.Contains(line, "Button B") {
            bX, bY = parseButtonValues(line)
         }

	 if strings.Contains(line, "Prize: X") {
	    x, y := parseButtonValues(line)
	    x += 10000000000000
	    y += 10000000000000

	    r, found := minClawTokens(x, y, aX, aY, bX, bY)

	    if found {
	        sum += r
	    }
	    fmt.Println("r:", r, "f:", found)
	 }

	 i++
    }

    return sum
}


func minClawTokens(x int, y int, aX int, aY int, bX int, bY int) (int, bool) {
	if x % bX == 0 && y % bY == 0 && x / bX == y / bY {
	    return x / bX, true
	}

	c := aX * y - aY * x
	fmt.Println("c: ", c)
	d := bY * aX - aY * bX
	fmt.Println("d: ", d)
	b := c / d
	a := (x - b * bX) / aX

	minTokens := a * 3 + b

	return minTokens, (a * aX + b * bX) == x && (a * aY + b * bY) == y
}

func parseButtonValues(line string) (int, int) {

	valueStr := strings.Split(line, ": ")[1]
	valuesRolled := strings.Split(valueStr, ", ")
	xStr := valuesRolled[0]
	yStr := valuesRolled[1]

	sep := "+"
	if !strings.Contains(xStr, sep) {
	    sep = "="
	}

	x, _ := strconv.Atoi(strings.Split(xStr, sep)[1])
	y, _ := strconv.Atoi(strings.Split(yStr, sep)[1])

	return x, y
}
