package problem_44

func isMatch(s string, p string) bool {
	states := make([][]bool, len(s)+1)
	for i := range states {
		states[i] = make([]bool, len(p)+1)
	}
	states[0][0] = true
	for j := 0; j < len(p); j++ {
		if p[j] == '*' {
			states[0][j+1] = states[0][j]
		}
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '?' {
				states[i+1][j+1] = states[i][j]
			} else if p[j] == '*' {
				states[i+1][j+1] = states[i][j+1] || states[i+1][j]
			}
		}
	}
	return states[len(s)][len(p)]
}
