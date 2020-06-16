package _31

import "log"

func rotate(matrix [][]int) {
	var newMatrix [][]int
	for i := 0; i < len(matrix); i++ {
		var item []int
		for j := len(matrix) - 1; j >= 0; j-- {
			item = append(item, matrix[j][i])
		}
		newMatrix = append(newMatrix, item)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[i][j] = newMatrix[i][j]
		}
	}
	log.Println(matrix)
}

func rotateBetter(matrix [][]int) {
	size := len(matrix)
	if 1 == size {
		return
	} else {
		for i := 0; i < size/2; i++ {
			move(matrix, i, size-1)
		}
	}
}

func move(matrix [][]int, s, l int) {
	i := s
	j := 0
	for k := s; k < l-s; k++ {
		j = k
		prev := matrix[i][j]
		i, j = j, l-i
		prev, matrix[i][j] = matrix[i][j], prev
		for s != i || k != j {
			i, j = j, l-i
			prev, matrix[i][j] = matrix[i][j], prev
		}
	}
}
