package problem_557

import (
	"bytes"
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	w := bytes.NewBuffer(make([]byte, 0, len(s)))
	si := 0
	for si < len(s) {
		if si != 0 {
			w.WriteByte(' ')
		}
		fi := si + 1
		for fi < len(s) {
			if s[fi] != ' ' {
				fi++
			} else {
				break
			}
		}
		for i := fi - 1; i >= si; i-- {
			w.WriteByte(s[i])
		}
		si = fi + 1
	}
	return w.String()
}
