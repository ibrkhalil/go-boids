package main

import (
	"fmt"
	"math/rand"
)

const (
	matrixSize = 250
)

var (
	matrixA = [matrixSize][matrixSize]int{
		{3, 1, -4},
		{2, -3, 1},
		{5, -2, 0},
	}
	matrixB = [matrixSize][matrixSize]int{
		{1, -2, -1},
		{0, 5, 4},
		{-1, -2, 3},
	}
	result = [matrixSize][matrixSize]int{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workOutRow(row int) {
	for col := 0; col < matrixSize; col++ {
		for i := 0; i < matrixSize; i++ {
			result[row][col] += matrixA[row][i] * matrixB[i][col]
		}
	}
}

func main() {
	generateRandomMatrix(&matrixA)
	generateRandomMatrix(&matrixB)
	for row := 0; row < matrixSize; row++ {
		workOutRow(row)
		fmt.Printf("%v\n", result[row])
	}
}