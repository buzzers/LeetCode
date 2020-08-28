package problem_657

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, c := range moves {
		if c == 'L' {
			x -= 1
		} else if c == 'R' {
			x += 1
		} else if c == 'U' {
			y += 1
		} else if c == 'D' {
			y -= 1
		}
	}
	return x == 0 && y == 0
}
