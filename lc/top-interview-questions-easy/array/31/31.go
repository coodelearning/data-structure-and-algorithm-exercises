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
		// 按层进行翻转
		for i := 0; i < size/2; i++ {
			moveForBetter(matrix, i, size-1)
		}
	}
}

func moveForBetter(matrix [][]int, s, l int) {
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

func rotateFromEn(matrix [][]int) {
	n := len(matrix) - 1
	for x := 0; x < len(matrix)/2; x++ {
		for y := x; y < n-x; y++ {
			// save top left
			t := matrix[x][y]
			// set top left to bottom left
			matrix[x][y] = matrix[n-y][x]
			// set bottom left to bottom right
			matrix[n-y][x] = matrix[n-x][n-y]
			// set bottom right to top right
			matrix[n-x][n-y] = matrix[y][n-x]
			// set top right to top left
			matrix[y][n-x] = t
		}
	}
}
