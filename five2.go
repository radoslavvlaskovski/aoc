package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func solveFive2() int {
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
		valid := false
		madeChanges := false
		split := strings.Split(line, ",")
		ordered := []int{}

		for _, v := range split {
		    x, _ := strconv.Atoi(v)
                    ordered = append(ordered, x)
	        }

		for !valid {
                    valid = true
                    set := make(map[int]int)
                    
		    fmt.Println("ord: ", ordered, )
		    for i, x := range ordered {
		        ix, exists := set[x]

		        if exists {
			    madeChanges = true
		            valid = false
			    ordered = addToIx(ordered, ix, x)
			    ordered = append(ordered[:i+1], ordered[i+2:]...)
			    break
		        } 

		        xRules, exists := rules[x]
		        if exists {
			    for _, r := range xRules {
				    _, ruleEx := set[r]
				    if ruleEx {
				        set[r] = min(i, set[r])
				    } else {
				        set[r] = i
				    }
			    }
		        }
		     }
		}

		if madeChanges {
	            fmt.Println("ord: ", ordered)
		    mid := ordered[len(ordered) / 2]
		    fmt.Println("mid: ", mid)
		    count += mid
		}
	}

    }

    return count;
}

func addToIx(arr []int, ix int, v int) []int {
    if ix == len(arr) {
        return append(arr, v)
    }

    result := append(arr[:ix+1], arr[ix:]...)
    result[ix] = v
    return result
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
