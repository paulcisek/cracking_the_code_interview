package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	tst := compress([]byte(os.Args[1]))
	fmt.Println(tst)
}

func compress(str []byte) string {
	var fullStr []byte
	for i := 0; i < len(str); i++ {
		currChar := str[i]
		counter := 1
		for j := i + 1; j < len(str); j++ {
			if currChar == str[j] {
				counter += 1
				i += 1
			}
		}
		fullStr = append(fullStr, str[i])
		fullStr = append(fullStr, []byte(strconv.Itoa(counter))...)
	}
	if len(fullStr) > len(str) {
		return string(str)
	} else {
		return string(fullStr)
	}
}
