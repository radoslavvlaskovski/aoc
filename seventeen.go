package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strconv"
    "strings"
    "reflect"
)

func solveSeventeen() int {
    file, err := os.Open("test")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    input := []string{}

    for scanner.Scan() {
         line := scanner.Text()
	 input = append(input, line)
    }

    instructions, r, flatI := parseInput17(input)
    stop := len(flatI)

    options := []int{}
    nOptions := []int{}
    for i := range 63 {
     nOptions = append(nOptions, i)
    }
    partialR := []int{}

    for len(flatI) > 0 {
	    partialR = append([]int{flatI[len(flatI) - 1]}, partialR...)
            flatI = flatI[:len(flatI) - 1]

	fmt.Println(len(nOptions), partialR)
	options = nOptions
	nOptions = []int{}
      for _, a := range options {
        out := []int{}
        r[0] = a
        i := 0

        for i < len(instructions) {
	    curr := instructions[i]
	    c := curr[0]
	    op := curr[1]
	    p, o := exec(c, op, &r)
	    if p != -1 {
	       i = p
	       continue
	    }
	    if o != -1 {
	        out = append(out, o)
	    }
	    i += 1
        }

        if reflect.DeepEqual(partialR, out) {
		if len(partialR) == stop {
		    return a
		}
		for i := a * 8; i <= a*8 + 8; i++ {
		    nOptions = append(nOptions, i)
		}
	}
    }
    }

    return 0
}

func parseInput17(lines []string) ([][]int, []int, []int) {
	instructions := [][]int{}
	flatI := []int{}
	r := []int{0, 0, 0}

	for i, line := range lines {
	    if i < 3 {
	       v, _ := strconv.Atoi(strings.Split(line, ": ")[1])
	       r[i] = v
	    }

	    if i == 4 {
	       l := strings.Split(line, ": ")[1]
	       values := strings.Split(l, ",")

	       for i := 0; i < len(values); i+= 2 {
		  inst, _ := strconv.Atoi(values[i])
		  op, _ := strconv.Atoi(values[i + 1])
		  flatI = append(flatI, inst, op)
	          instructions = append(instructions, []int{inst, op})
	       }
	    }
	}

	return instructions, r, flatI
}

func exec(c int, op int, r *[]int) (int, int) {
    p := -1
    o := -1
    if c == 0 || c == 6 || c == 7 {
        res := adv(op, r)
	if c == 0 {
	    (*r)[0] = res
	} else {
	    (*r)[c % 5] = res
	}
    }
    if c == 1 {
	(*r)[1] = (*r)[1] ^ op
    }
    if c == 2 {
         (*r)[1] = getOp(op, *r) & 7
    }
    if c == 3 && (*r)[0] != 0 {
        p = op
    }
    if c == 4 {
	(*r)[1] = (*r)[1] ^ (*r)[2]
    }
    if c == 5 {
        o = getOp(op, *r) & 7
    }

    return p, o
}

func adv(op int, r *[]int) int {
	return (*r)[0] >> getOp(op, *r)
}

func getOp(op int, r []int) int {
    if op < 4 {
        return op
    }

    return r[op % 4]
}
