package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
    "strconv"
)

type evolution struct {
    v string
    e int
}

func solveEleven2() int {
	file, err := os.Open("test")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    blinks := 50
    var values []string
    count := 0

    for scanner.Scan() {
         line := scanner.Text()

	 values = strings.Split(line, " ")
	 break
    }

    fmt.Println(values)

    evs := []evolution{evolution{v: values[0], e: 0}}

    for len(evs) != 0 {
	    ev := evs[0]
	    evs = evs[1:]

	    if ev.e == blinks {
	       count++

	       continue
	    }

	    for _, x := range executeEvolution2(ev.v) {
		    evs = append(evs, evolution{v: x, e: ev.e + 1})
	    }

    }

    return count
}

func executeEvolution2(x string) []string {
	newValues := []string{}

		if x == "0" {
		   newValues = append(newValues, "1")
		   return newValues
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
			return newValues
		}

		k, _ := strconv.Atoi(x)

		newValues = append(newValues, strconv.Itoa(k * 2024))

	return newValues
}

