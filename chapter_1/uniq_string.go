package main

import (
	"fmt"
	"os"
)

func uniqString(strToBeChecked string) bool {
	stringLength := len(strToBeChecked)
	for i := 0; i < stringLength; i++ {
		testedChar := strToBeChecked[i]
		for j := i + 1; j < stringLength; j++ {
			if (testedChar == strToBeChecked[j]) && (i != j) {
				return false
			}
		}
	}
	return true
}

func main() {
	stringToBeChecked := os.Args[1]
	uniq := uniqString(stringToBeChecked)
	fmt.Println(uniq)
}
