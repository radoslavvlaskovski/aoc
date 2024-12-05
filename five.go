package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func solveFive() int {
	file, err := os.Open("input5")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    count := 0
    scanner := bufio.NewScanner(file)

    rules := make(map[int][]int)

    for scanner.Scan() {
        line := scanner.Text()
        
	if strings.Contains(line, "|") {
		split := strings.Split(line, "|")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		v, exists := rules[y]
		if exists {
		    rules[y] = append(v, x)
		} else {
		   rules[y] = []int{x}
		}
	}

	if strings.Contains(line, ",") {
		valid := true
		split := strings.Split(line, ",")
                set := make(map[int]bool)

		for _, v := range split {
		    x, _ := strconv.Atoi(v)

		    _, exists := set[x]

		    if exists {
		        valid = false
			break
		    }

		    rules, exists := rules[x]
		    if exists {
			    for _, r := range rules {
			       set[r] = true
			    }
		    }
		}

		if valid {
		    mid, _ := strconv.Atoi(split[len(split) / 2])
		    fmt.Println("mid: ", mid)
		    count += mid
		}
	}

    }

    return count;
}

