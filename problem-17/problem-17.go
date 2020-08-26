package problem_17

var numbers = [][]string{{}, {}, {"a", "b", "c"}, {"d", "e", "f"},
	{"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"},
	{"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}

func letterCombinations(digits string) []string {
	// 0 个直接返回空
	if len(digits) == 0 {
		return []string{}
	}
	// 1 个直接返回对应的字符表
	if len(digits) == 1 {
		return numbers[digits[0]-'0']
	}
	// 递归 digits[1:] 的解
	combinations0 := letterCombinations(digits[1:])
	combinations := make([]string, 0, len(combinations0)*len(numbers[digits[0]-'0'])+len(numbers[digits[0]-'0']))
	// 组合 digits[0] 的字符表和 digits[1:] 的解得到 digits 的解
	for _, l := range numbers[digits[0]-'0'] {
		for _, r := range combinations0 {
			combinations = append(combinations, l+r)
		}
	}

	return combinations
}
