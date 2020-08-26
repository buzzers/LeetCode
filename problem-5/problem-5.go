package problem_5

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	// states[i][j] 代表 s[i:j+1] 是否是回文
	states := make([][]bool, len(s))
	for i := range states {
		states[i] = make([]bool, len(s))
	}

	// left, right 表示当前找到的最长回文串的索引
	left, right := 0, 0
	// l 表示当前搜索的区间长度
	for l := 1; l < len(s); l++ {
		// i 表示当前搜索的区间起始索引
		for i := 0; i+l < len(s); i++ {
			if l < 3 {
				// 如果当前搜索的长度小于 3，那么直接看区间的首位是否相等即可判断是否是回文串。
				// 即 x0x1x2，只要 x0 == x2，就是回文。
				states[i][i+l] = s[i] == s[i+l]
			} else {
				// 否则当首尾相等且去掉首尾后任然是回文串，即是回文。
				states[i][i+l] = s[i] == s[i+l] && states[i+1][i+l-1]
			}
			// 记录最长回文
			if states[i][i+l] && (l > right-left) {
				left, right = i, i+l
			}
		}
	}
	return s[left : right+1]
}
