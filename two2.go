package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "strconv"
    "fmt"
)

func solveTwo2() int {
	file, err := os.Open("input2")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    count := 0
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
	
	words := strings.Split(line, " ")
        
       intList := make([]int, 0, len(words))

       for _, str := range words {
           num, _ := strconv.Atoi(str)
           intList = append(intList, num)
       }

       fmt.Println(intList)

       if safetyCheck2(intList, false) {
           count++
       } else {
           reverse(intList)
	   if safetyCheck2(intList, false) {
	       count++;
	   }
       }
    }

    return count;
}


func safetyCheck2(data []int, dUsed bool) bool {
    if len(data) == 1 {
        return true
    }

    for i := 1; i < len(data); i++ {

	if !inLim(data[i], data[i - 1]) {
	    if dUsed {
	        return false
	    }

	    if i == len(data) - 1 {
	         return true
	    }

	    var removeLeft bool

	    if i - 1 == 0 || inLim(data[i], data[i - 2]) {
                removeLeft = safetyCheck2(data[i:], true)
	    }

	    removeRight := safetyCheck2(append([]int{data[i - 1]}, data[i+1:]...), true)

	    return removeLeft || removeRight
	}

    }

    return true
}

func reverse(arr []int) {
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}

