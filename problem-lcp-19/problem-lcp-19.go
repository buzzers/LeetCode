package problem_lcp_19

import "math"

func minimumOperations(leaves string) int {
	state := make([][3]int, len(leaves))
	state[0][0] = isRed(leaves[0])
	state[0][1] = math.MaxInt32
	state[0][2] = math.MaxInt32
	for i := 1; i < len(leaves); i++ {
		state[i][0] = state[i-1][0] + isRed(leaves[i])
		state[i][1] = min(state[i-1][0], state[i-1][1]) + isYellow(leaves[i])
		state[i][2] = min(state[i-1][1], state[i-1][2]) + isRed(leaves[i])
	}
	return state[len(leaves)-1][2]
}

func isYellow(c byte) int {
	if c == 'y' {
		return 0
	}
	return 1
}

func isRed(c byte) int {
	if c == 'r' {
		return 0
	}
	return 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
