package main

import (
	"fmt"
	"os"
)

func reverse(inputString string) string {
	stringToBeReversed := []byte(inputString)
	stringLength := len(stringToBeReversed) - 1
	for i := 0; i <= stringLength/2; i++ {
		waitingChar := stringToBeReversed[i]
		stringToBeReversed[i] = stringToBeReversed[stringLength-i]
		stringToBeReversed[stringLength-i] = waitingChar
	}

	return string(stringToBeReversed)

}

func main() {
	stringToBeReverted := os.Args[1]
	fmt.Println(reverse(stringToBeReverted))
}
