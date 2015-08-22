package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := []byte(os.Args[1])
	strTrueLength, _ := strconv.Atoi(os.Args[2])
	strLength := len(str)
	fmt.Println(changeSpaces(str, strTrueLength, strLength))
}

func moveArrayElementsByOne(str []byte, strLength, startingIndex int) []byte {
	for i := strLength - 1; i > startingIndex; i-- {
		str[i] = str[i-1]
	}
	return str
}

func moveArrayElementsByThree(str []byte, strLength, startingIndex int) []byte {
	for j := 0; j < 2; j++ {
		str = moveArrayElementsByOne(str, strLength, startingIndex)
	}
	return str
}

func setNewValuesForSpace(str []byte, startingIndex int) []byte {
	str[startingIndex] = '%'
	str[startingIndex+1] = '2'
	str[startingIndex+2] = '0'
	return str
}

func changeSpaces(str []byte, trueLength, strLength int) string {
	if trueLength == 1 {
		return string(setNewValuesForSpace(str, 0))
	}
	for i := 0; i < trueLength; i++ {
		if str[i] == ' ' {
			str = moveArrayElementsByThree(str, strLength, i)
			str = setNewValuesForSpace(str, i)
			trueLength = trueLength + 2
		}
	}
	return string(str[:])
}
