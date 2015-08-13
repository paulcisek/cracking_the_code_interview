package main

import (
	"fmt"
	"os"
)

func main() {
	str1 := os.Args[1]
	str2 := os.Args[2]
	fmt.Println(isPermutation(str1, str2))
}

func isPermutation(string1, string2 string) bool {
	stringMap := make(map[rune]bool)
	for _, element := range string1 {
		stringMap[element] = true
	}
	for _, element := range string2 {
		if !stringMap[element] {
			return false
		}
	}
	return true
}
