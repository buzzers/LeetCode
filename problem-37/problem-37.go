package problem_37

func solveSudoku(board [][]byte) {
	rows, cols, box, blanks := buildMaps(board)
	solveSudokuInternal(board, rows, cols, box, blanks)
}

func solveSudokuInternal(board [][]byte, rows, cols, box []map[byte]bool, blanks [][2]int) bool {
	if len(blanks) == 0 {
		return true
	}
	y, x := blanks[0][0], blanks[0][1]
	for i := 1; i <= 9; i++ {
		val := byte(i + '0')
		if !rows[y][val] && !cols[x][val] && !box[3*(y/3)+x/3][val] {
			board[y][x] = val
			updateMaps(val, x, y, rows, cols, box, true)
			if solveSudokuInternal(board, rows, cols, box, blanks[1:]) {
				return true
			} else {
				updateMaps(val, x, y, rows, cols, box, false)
			}
		}
	}
	board[y][x] = '.'
	return false
}

func updateMaps(val byte, x, y int, rows, cols, box []map[byte]bool, put bool) {
	rows[y][val] = put
	cols[x][val] = put
	box[3*(y/3)+x/3][val] = put
}

func buildMaps(board [][]byte) ([]map[byte]bool, []map[byte]bool, []map[byte]bool, [][2]int) {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	box := make([]map[byte]bool, 9)
	blanks := make([][2]int, 0, 81)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] != '.' {
				rows[y][board[y][x]] = true
				cols[x][board[y][x]] = true
				box[3*(y/3)+x/3][board[y][x]] = true
			} else {
				blanks = append(blanks, [2]int{y, x})
			}
		}
	}
	return rows, cols, box, blanks
}
