package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func solveFour2() int {
	file, err := os.Open("input4")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    count := 0
    scanner := bufio.NewScanner(file)

    input := []string{}

    for scanner.Scan() {
        line := scanner.Text()
	fmt.Println("processing line")

	input = append(input, line)
    }

    var y = len(input)
    var x = len(input[0])

    fmt.Println("x: ", x, " y: ", y)

    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {
	    if string(input[i][j]) != "A" {
	        continue
	    }
	    left := check2(input, []int{i + 1, j +1}, []int{i - 1, j - 1}) || check2(input, []int{i - 1, j - 1}, []int{i + 1, j + 1})

	    right := check2(input, []int{i - 1, j + 1}, []int{i + 1, j - 1}) || check2(input, []int{i + 1, j - 1}, []int{i - 1, j + 1})

	    if left && right {
                count++
	    }
    }

    }

    return count;
}

func check2(input []string, m []int, s []int) bool {
	y := len(input)
	x := len(input[0])

	return checkLetter(x, y, input, m, "M") && checkLetter(x, y, input, s, "S")
}

func checkLetter2(x int, y int, input []string, c []int, l string) bool {
    if c[0] < 0 || c[0] >= y || c[1] < 0 || c[1] >= x {
         return false
    }

    return string(input[c[0]][c[1]]) == l
}


