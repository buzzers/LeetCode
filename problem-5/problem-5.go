package problem_5

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	states := make([][]bool, len(s))
	for i := range states {
		states[i] = make([]bool, len(s))
	}

	left, right := 0, 0
	for l := 1; l < len(s); l++ {
		for i := 0; i+l < len(s); i++ {
			if l < 3 {
				states[i][i+l] = s[i] == s[i+l]
			} else {
				states[i][i+l] = s[i] == s[i+l] && states[i+1][i+l-1]
			}
			if states[i][i+l] && (l > right-left) {
				left, right = i, i+l
			}
		}
	}
	return s[left : right+1]
}
