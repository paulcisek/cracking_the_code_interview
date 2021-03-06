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
	if len(string1) != len(string2) {
		return false
	}
	stringMap := make(map[rune]int)
	for _, element := range string1 {
		stringMap[element] = stringMap[element] + 1
	}
	for _, element := range string2 {
		if stringMap[element] {
			if stringMap[element] == 0 {
				return false
			}
			stringMap[element] = stringMap[element] - 1
		} else {
			return false
		}
	}
	return true
}
