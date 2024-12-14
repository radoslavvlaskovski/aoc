package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

type fileId struct {
    id int
    count int
}

func solveNine() int {
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
	     v, _ := byteToDigit(byte(x))

	     if i % 2 != 0 {
		     for range v {
			     last := ids[len(ids) - 1]
			     out += fmt.Sprintf("%d", last.id)
			     sum += last.id * curr
			     curr++
			     if last.count == 1 {
		                 ids = ids[:len(ids) - 1]
			     } else {
			        ids[len(ids) - 1].count--
			     }
			     if len(ids) == 0 {
				     fmt.Println(out)
                                return sum
                             }
		     }
	     } else {
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

func byteToDigit(b byte) (int, error) {
    if b >= '0' && b <= '9' {
        return int(b - '0'), nil // Subtract '0' to get the numeric value
    }
    return 0, fmt.Errorf("byte %q is not a digit", b)
}
