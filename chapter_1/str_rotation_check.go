package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strRotationContains(os.Args[1], os.Args[2]))
}

func strRotationContains(strRotation, subStr string) bool {
	if len(strRotation) != len(subStr) {
		return false
	}
	fullString := generateFullStringFromRotatedString(strRotation)
	return isSubString(fullString, subStr)

}

func isSubString(str, subStr string) bool {
	return strings.Contains(str, subStr)
}

func generateFullStringFromRotatedString(rotatedString string) string {
	var buffer bytes.Buffer
	for i := 0; i < 2; i++ {
		buffer.WriteString(rotatedString)
	}
	return buffer.String()
}
