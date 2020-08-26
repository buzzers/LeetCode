package problem_44

func isMatch(s string, p string) bool {
	// states[i][j] 代表 s 的前 i 个字符能否被 p 的前 j 个字符匹配
	states := make([][]bool, len(s)+1)
	for i := range states {
		states[i] = make([]bool, len(p)+1)
	}

	// 下述的 i 和 j 都是指 s 和 p 的索引，因此在访问 states 时需要加 1。

	// states[0][0] 即 s 的前 0 个字符能否被 p 的前 0 个字符匹配
	// 必定为 true
	states[0][0] = true
	// states[0][j+1] 即 s 的前 0 个字符能否被 p 的前 j 个字符匹配
	// 此处处理 s 为空且 p 仅包含 "*" 模式
	for j := 0; j < len(p); j++ {
		if p[j] == '*' {
			// 如果 p[j] 是 '*'，
			// 那么 s 的前 0 个字符能否被 p 的前 j+1 个字符匹配，
			// 就看 s 的前 0 个字符能否被 p 的前 j 个字符匹配即可。
			// 即 p[j] 的模式 "*" 匹配 0 个字符
			states[0][j+1] = states[0][j]
		}
	}
	// 状态转移方程
	// 如果 s[i] 直接匹配 p[j] 或者模式串为 '?'，则 states[i+1][j+1] 就看 states[i][j]；
	// 如果 p[j] 为 '*'，则 states[i+1][j+1] 就看 states[i][j+1] 或 states[i+1][j]。
	// 解释
	// 如果 s[i] == p[j] 或者 p[j] == '?'，则表示当前的字符和模式匹配，那么前 i+1 个字符是否被前 j+1 个模式匹配，
	// 直接看前面的是否匹配即可，即 states[i+1][j+1] = states[i][j]。
	// 如果 p[j] == '*'，则两种情况的任意一种匹配即可：
	// 当前的模式 '*' 消耗 s[i] 和不消耗 s[i]，即至少匹配一个和匹配 0 个。
	//   states[i][j+1] (s 的前 i 个字符能否被 p 的前 j+1 个字符匹配，即 p[j] 的模式串 '*' 消耗掉 s[i] 但自身不被消耗)
	//   states[i+1][j] (s 的前 i+1 个字符能否被 p 的前 j 个字符匹配，即 p[j] 的模式串 '*' 不匹配 s[i] 且消耗自身)
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
