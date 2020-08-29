package problem_214

import "bytes"

func shortestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	// 从起始位置寻找最长的回文子串
	l := 1
	for i := 2; i <= len(s); i++ {
		isPalindrome := true
		for j := 0; j < i/2; j++ {
			if s[j] != s[i-j-1] {
				isPalindrome = false
				break
			}
		}
		if isPalindrome {
			l = i
		}
	}

	// 解即为 反转的最长回文子串剩余部分 + 原字符串
	w := bytes.NewBuffer(make([]byte, 0, 2*len(s)-l))
	for i := len(s) - 1; i >= l; i-- {
		w.WriteByte(s[i])
	}
	w.WriteString(s)
	return w.String()
}
