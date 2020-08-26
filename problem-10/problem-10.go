package problem_10

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
	// states[0][j+1] 即 s 的前 0 个字符能否被 p 的前 j+1 个字符匹配
	// 此处处理 s 为空且 p 仅包含 "a*" 模式
	for j := 1; j < len(p); j++ {
		if p[j] == '*' {
			// 如果 p[j] 是 '*'，
			// 那么 s 的前 0 个字符能否被 p 的前 j+1 个字符匹配，
			// 就看 s 的前 0 个字符能否被 p 的前 j-1 个字符匹配即可。
			// 即 p[j:j+1] 的模式 "a*" 匹配 0 个字符
			states[0][j+1] = states[0][j-1]
		}
	}
	// 状态转移方程
	// 如果 s[i] 直接匹配 p[j] 或者模式串为 '.'，则 states[i+1][j+1] 就看 states[i][j]；
	// 如果 p[j] 为 '*'，则需要分两个情况：
	// 如果 s[i] 和 p[j-1] 匹配或者 p[j-1] 为 '.'，则 states[i+1][j+1] 就看 states[i][j+1] 或 states[i+1][j] 或 states[i+1][j-1]，
	// 否则 states[i+1][j+1] 就看 states[i+1][j-1]。
	// 解释
	// 如果 s[i] == p[j] 或者 p[j] == '.'，则表示当前的字符和模式匹配，那么前 i+1 个字符是否被前 j+1 个模式匹配，
	// 直接看前面的是否匹配即可，即 states[i+1][j+1] = states[i][j]。
	// 如果 p[j] == '*'，情况较为复杂：
	// 如果 s[i] == p[j-1] || p[j-1] == '.'，则表示当前字符能被当前的 "a*" 模式串(即 p 的前一个字符)匹配，
	// 那么当以下三种情况任意一种匹配时，都可以配成功：
	// 当前的 "a*" 模式匹配 0 个字符、当前的 "a*" 模式匹配 1 个字符和当前的 "a*" 模式匹配不止一个字符。
	// 这对应的状态是：
	//   states[i+1][j-1] (s 的前 i+1 个字符能否被 p 的前 j-1 个字符匹配，即 p[j:j+1] 的模式串不生效)
	//   states[i+1][j] (s 的前 i+1 个字符能否被 p 的前 j 个字符匹配，即 p[j:j+1] 的模式串仅看 "a*" 的 'a' 部分，忽略 '*')
	//   states[i][j+1] (s 的前 i 个字符能否被 p 的前 j+1 个字符匹配，即 p[j:j+1] 的模式串将当前 s 的字符消耗但模式串本身不消耗)
	// 而如果当前字符不能被当前的 "a*" 模式串匹配，则只能选择 states[i+1][j-1]，即匹配 0 个字符。
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '.' {
				states[i+1][j+1] = states[i][j]
			} else if p[j] == '*' {
				if s[i] == p[j-1] || p[j-1] == '.' {
					states[i+1][j+1] = states[i][j+1] || states[i+1][j] || states[i+1][j-1]
				} else {
					states[i+1][j+1] = states[i+1][j-1]
				}
			}
		}
	}
	return states[len(s)][len(p)]
}
