package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
)

func solveNine2() int {
	file, err := os.Open("input9")
	fmt.Println("solving...")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    sum := 0
    
    for scanner.Scan() {
         line := scanner.Text()

	 ids := []fileId{}

	 for i, x := range line {
	      if i % 2 != 0 {
	          continue
	      }

	      id := i / 2
	      idCount, _ := byteToDigit(byte(x))
	      f := fileId{id: id, count: idCount}
	      ids = append(ids, f)
	 }

	 fmt.Println(ids)
	 out := ""

	 curr := 0

	 for i, x := range line {
             if len(ids) == 0 {
                  fmt.Println(out)
                  return sum
             }
	     v, _ := byteToDigit(byte(x))

	     if i % 2 != 0 {
		     next, found := findAndPop(&ids, v)

		     rest := v

		     for found {
			 for range next.count {
			     out += fmt.Sprintf("%d", next.id)
			     sum += next.id * curr
			     curr++
			     rest--
			 }

			 next, found = findAndPop(&ids, rest)
		     }

		     curr += rest
		     out += strings.Repeat(".", rest)
	     } else {
		     if i / 2 < ids[0].id {
		         curr += v
			 out += strings.Repeat(".", v)
			 continue
		     }
		     first := ids[0]

		     for range first.count {
		          sum += first.id * curr
			  out += fmt.Sprintf("%d", first.id)
			  curr++
		     }

		     if len(ids) == 1 {
			     fmt.Println(out)
                          return sum
                     }
		     ids = ids[1:]
	     }
	 }

	 break
    }

    return sum
}

func findAndPop(fileIds *[]fileId, size int) (fileId, bool) {
	i := len(*fileIds) - 1

	for i >= 0 {
	    if (*fileIds)[i].count <= size {
		res := (*fileIds)[i]
		*fileIds = append((*fileIds)[:i], (*fileIds)[i+1:]...)

		return res, true
	    }
	    i--
	}

	return fileId{}, false
}

