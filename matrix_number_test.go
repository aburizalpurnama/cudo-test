package main

import (
	"fmt"
	"testing"
)

func TestMatrixNumber(t *testing.T) {
	result := GetRequiredNumber(5, 2, 4)
	fmt.Printf("result: %v\n", result)
}

func GetRequiredNumber(size, x, y int) int {
	if x < 1 || y < 1 {
		return 0
	}

	max := (size + 1) * size / 2

	numbers := generateNumbers(max)

	matrixNumber := generateMatrix(size, numbers)

	fmt.Printf("matrixNumber: %v\n", matrixNumber)

	return matrixNumber[size-y][x-1]
}

func generateMatrix(size int, numbers []int) [][]int {
	matrixNumber := make([][]int, size)
	for i := range matrixNumber {
		matrixNumber[i] = make([]int, i+1)
	}

	substactor := 0
	for i := 0; i < size; i++ {
		index := len(numbers) - i - 1 - substactor
		for j := 0; j < i+1; j++ {
			matrixNumber[i][j] = numbers[index]
			index++
			substactor++
		}
	}
	return matrixNumber
}

func generateNumbers(max int) []int {
	var numbers []int
	var number int
	for i := 0; i < max; i++ {
		number = number + (i + 1)
		numbers = append(numbers, number)
	}
	return numbers
}
