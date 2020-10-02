package problem_771

func numJewelsInStones(J string, S string) int {
	bitmap := make(map[byte]struct{}, len(J))
	for _, b := range J {
		bitmap[byte(b)] = struct{}{}
	}
	n := 0
	for _, b := range S {
		if _, ok := bitmap[byte(b)]; ok {
			n++
		}
	}
	return n
}
