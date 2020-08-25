package problem_22

func generateParenthesis(n int) []string {
	states := [][]string{{""}}
	for i := 1; i <= n; i++ {
		var state []string
		for j := 0; j < i; j++ {
			for _, l := range states[j] {
				for _, r := range states[i-j-1] {
					state = append(state, "("+l+")"+r)
				}
			}
		}
		states = append(states, state)
	}
	return states[n]
}
