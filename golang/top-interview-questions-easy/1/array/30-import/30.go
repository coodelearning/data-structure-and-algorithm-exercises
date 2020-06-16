package _30

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		mapRow := map[byte]bool{}
		mapCol := map[byte]bool{}
		mapPart := map[byte]bool{}
		for j := 0; j < 9; j++ {
			// 判断行重复
			if board[j][i] != '.' && mapRow[board[j][i]] {
				return false
			}
			mapRow[board[j][i]] = true
			// 判断列重复户
			if board[i][j] != '.' && mapCol[board[i][j]] {
				return false
			}
			mapCol[board[i][j]] = true
			// 判断小区域是否有重复
			colIndex := (i%3)*3 + j%3
			rowIndex := (i/3)*3 + j/3
			if board[colIndex][rowIndex] != '.' && mapPart[board[colIndex][rowIndex]] {
				return false
			}
			mapPart[board[colIndex][rowIndex]] = true
		}
	}
	return true
}

func BetterIsValidSudoku2(board [][]byte) bool {
	// 行列检测
	for i := 0; i < 9; i++ {
		mp1 := map[byte]bool{}
		mp2 := map[byte]bool{}
		mp3 := map[byte]bool{}
		for j := 0; j < 9; j++ {
			// row
			if board[i][j] != '.' {
				if mp1[board[i][j]] {
					return false
				}
				mp1[board[i][j]] = true
			}
			// column
			if board[j][i] != '.' {
				if mp2[board[j][i]] {
					return false
				}
				mp2[board[j][i]] = true
			}
			// part
			row := (i%3)*3 + j%3
			col := (i/3)*3 + j/3
			if board[row][col] != '.' {
				if mp3[board[row][col]] {
					return false
				}
				mp3[board[row][col]] = true
			}

		}
	}
	return true
}

// leetcode 题解
// Ps.噢,这该死的位运算
func betterIsValidSudoku(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')                      // 当前数字的 二进制数位 位置
			bi := i/3 + j/3*3                                   // 3x3的块索引号
			if (row[i]&cur)|(col[j]&cur)|(block[bi]&cur) != 0 { // 使用与运算查重
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}
