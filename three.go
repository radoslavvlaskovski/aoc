package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "unicode"
    "strings"
    "strconv"
)

func solveThree() int {
	file, err := os.Open("input3")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    count := 0
    scanner := bufio.NewScanner(file)

    enabled := true

    for scanner.Scan() {
        line := scanner.Text()
	fmt.Println("processing line")

	for i := 4; i < len(line); i++ {
	    if i > 7 && line[i-7:i] == "don't()" {
	    	enabled = false
	    }
            if line[i-4:i] == "do()" {
	        enabled = true
	    }

	    if !enabled {
		continue
	    }

	    if line[i - 4:i] == "mul(" {
	        for j := i; j < len(line); j++ {
		     charStr := string(line[j])
		     if isDigit(charStr) || charStr == "," {
			  continue
		     }
		     if charStr == ")" {
		          count += getResult(line[i:j])
			  fmt.Println(line[i:j])
		     }
		     break
		}
	    }
        }
    }

    return count;
}

func getResult(f string) int {
    if len(f) > 7 {
        return 0
    }

    words := strings.Split(f, ",")
    if len(words) != 2 {
        return 0
    }
    left := words[0]
    right := words[1]

    if len(left) == 0 || len(right) == 0 || len(left) > 3 || len(right) > 3 || string(left[0]) == "0" || string(right[0]) == "0" {
        return 0
    }

    leftValue, _ := strconv.Atoi(left)
    rightValue, _ := strconv.Atoi(right)
    
    return leftValue * rightValue
}

func isDigit(s string) bool {
    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}
