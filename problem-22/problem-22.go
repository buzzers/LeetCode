package problem_22

func generateParenthesis(n int) []string {
	// states[i] 代表 i 对括号的解
	states := [][]string{{""}}
	// 状态转移方程
	// states[i] = '(' + states[j] + ')' + states[i-j-1]
	// 其中 j < i
	// 解释
	// i 对括号的解，可以从 i - 1 对括号的解得到
	// 每个解一定从一个 '(' 开始，并有一个和它配对的 ')'
	// 于是需要寻找所有可能的 ')‘ 的位置
	// 即计算集合 '(' + [j 对括号的解] + ')' + [i-j-1 对括号的解]
	// 可枚举出所有 ')' 的位置，便得到 i 对括号的解
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
