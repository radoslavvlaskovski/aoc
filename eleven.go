package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func solveEleven() int {
	file, err := os.Open("test")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    blinks := 75
    var values []string

    for scanner.Scan() {
         line := scanner.Text()

	 values = strings.Split(line, " ")
	 break
    }

    fmt.Println(values)

    for i := range blinks {
	    fmt.Println("blink: ", i)
        values = executeEvolution(values)
    }

    return len(values)
}

func executeEvolution(values []string) []string {
	newValues := []string{}

	for _, x := range values {

		if x == "0" {
		   newValues = append(newValues, "1")
		   continue
		}

		if len(x) % 2 == 0 {
			left := x[:len(x) / 2]

			rightStart := len(x) / 2

			for _, c := range x[len(x) / 2:] {
			    if byte(c) != byte('0') {
			       break
			    }

			    rightStart++
			}

			if rightStart == len(x) {
			    rightStart--
			}

			newValues = append(newValues, left, x[rightStart:])
			continue
		}

		k, _ := strconv.Atoi(x)

		newValues = append(newValues, strconv.Itoa(k * 2024))
	}

	return newValues
}

