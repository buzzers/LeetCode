package problem_51

import "strings"

func solveNQueens(n int) [][]string {
	// board[i] 代表第 i 行的皇后所在列
	board := make([]int, n)
	for i := range board {
		// 未放置值为 -1
		board[i] = -1
	}
	return solveNQueensInternal(n, board, 0)
}

func solveNQueensInternal(n int, board []int, row int) [][]string {
	var results [][]string
	// 枚举每列是否可放置
	for col := 0; col < n; col++ {
		if isCanPlace(board, row, col) {
			// 放置皇后
			board[row] = col
			if row == n-1 {
				// 如果已经放满则输出一个结果
				results = append(results, buildBoard(board))
			} else {
				// 如果未放满则继续递归求解下一行
				results = append(results, solveNQueensInternal(n, board, row+1)...)
			}
			// 取消放置
			board[row] = -1
		}
	}
	return results
}

func buildBoard(board []int) []string {
	result := make([]string, len(board))
	for row, col := range board {
		result[row] = strings.Repeat(".", col) + "Q" + strings.Repeat(".", len(board)-col-1)
	}
	return result
}

func isCanPlace(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || abs(i-row) == abs(board[i]-col) {
			return false
		}
	}
	return true
}

func abs(val int) int {
	if val < 0 {
		val *= -1
	}
	return val
}
