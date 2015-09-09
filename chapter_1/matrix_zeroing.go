package main

import (
	"fmt"
)

type Matrix [][]int

func main() {
	matrix := Matrix{
		[]int{0, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 0, 4, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 0, 5},
	}
	fmt.Println(zeroing(matrix))
}

func zeroing(matrix Matrix) Matrix {
	zeroedRows := make(map[int]bool)
	zeroedCols := make(map[int]bool)
	for i, item := range matrix {
		for j, _ := range item {
			_, rowZeroed := zeroedRows[i]
			_, colZeroed := zeroedCols[j]
			if !rowZeroed && !colZeroed && matrix[i][j] == 0 {
				zeroedRows[i] = true
				zeroedCols[j] = true
			}
		}
	}
	zeroRows(&matrix, zeroedRows)
	zeroCols(&matrix, zeroedCols)

	return matrix
}

func zeroRows(matrix *Matrix, zeroedRows map[int]bool) {
	for key, _ := range zeroedRows {
		for i := 0; i < len((*matrix)[key]); i++ {
			(*matrix)[key][i] = 0
		}
	}
}

func zeroCols(matrix *Matrix, zeroedCols map[int]bool) {
	for key, _ := range zeroedCols {
		for i := 0; i < len(*matrix); i++ {
			(*matrix)[i][key] = 0
		}
	}
}
