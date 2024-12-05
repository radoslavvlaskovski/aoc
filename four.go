package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func solveFour() int {
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

    v := 0
    vb := 0
    h := 0
    hb := 0

    fmt.Println("x: ", x, " y: ", y)

    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {
	    if string(input[i][j]) != "X" {
	        continue
	    }
            h += check(input, []int{i + 1, j}, []int{i + 2, j}, []int{i + 3, j}) 
	    hb += check(input, []int{i - 1, j}, []int{i - 2, j}, []int{i - 3, j})
	    v += check(input, []int{i, j + 1}, []int{i, j + 2}, []int{i, j + 3})
	    vb += check(input, []int{i, j - 1}, []int{i, j - 2}, []int{i, j - 3})
	    count += check(input, []int{i - 1, j - 1}, []int{i - 2, j - 2}, []int{i - 3, j - 3})
	    count += check(input, []int{i + 1, j + 1}, []int{i + 2, j + 2}, []int{i + 3, j + 3})
	    count += check(input, []int{i + 1, j - 1}, []int{i + 2, j - 2}, []int{i + 3, j - 3})
	    count += check(input, []int{i - 1, j + 1}, []int{i - 2, j + 2}, []int{i - 3, j + 3})
    }

    }

    fmt.Println("v: ", v, "vb: ", vb, "h: ", h, "hb: ", hb)

    return count + v + vb + h + hb;
}

func check(input []string, m []int, a []int, s []int) int {
	y := len(input)
	x := len(input[0])

	if checkLetter(x, y, input, m, "M") && checkLetter(x, y, input, a, "A") && checkLetter(x, y, input, s, "S") {
	    return 1
	}
	
	return 0
}

func checkLetter(x int, y int, input []string, c []int, l string) bool {
    if c[0] < 0 || c[0] >= y || c[1] < 0 || c[1] >= x {
         return false
    }

    return string(input[c[0]][c[1]]) == l
}


