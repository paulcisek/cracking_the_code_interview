package main

import (
	"fmt"
)

type Matrix [][]int

func main() {
	matrix := Matrix{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4},
	}
	fmt.Println(rotate(matrix))

}

func rotate(matrix Matrix) Matrix {
	for i, item := range matrix {
		for j, _ := range item {
			if j > i {
				tmp := matrix[i][j]
				matrix[i][j] = matrix[j][i]
				matrix[j][i] = tmp
			}
		}
	}
	return matrix
}
