package problem_79

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for y, row := range board {
		for x := range row {
			if existInternal(board, word, visited, x, y, 0) {
				return true
			}
		}
	}
	return false
}

func existInternal(board [][]byte, word string, visited [][]bool, x, y, k int) bool {
	if board[y][x] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	visited[y][x] = true
	defer func() { visited[y][x] = false }()
	for _, direction := range directions {
		x0, y0 := x+direction[0], y+direction[1]
		if 0 <= x0 && x0 < len(board[0]) && 0 <= y0 && y0 < len(board) && !visited[y0][x0] {
			if existInternal(board, word, visited, x0, y0, k+1) {
				return true
			}
		}
	}
	return false
}
