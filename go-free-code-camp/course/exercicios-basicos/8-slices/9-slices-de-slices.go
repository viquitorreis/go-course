package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matriz := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0) // precisamos de um slice simples para representar as linhas
		for j := 0; j < cols; j++ {
			row = append(row, i*j) // cada célula precisa ser o resultado de i*j
		}

		matriz = append(matriz, row) // precisamos fazer o append da linha na matriz
	}

	return matriz
}

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
