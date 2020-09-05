package problem_60

import "fmt"

var (
	factorials = []int{1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
)

func getPermutation(n int, k int) string {
	values := make([]int, n)
	for i := range values {
		values[i] = i + 1
	}
	return getPermutationInternal(values, k)
}

func getPermutationInternal(values []int, k int) string {
	if len(values) == 1 {
		return fmt.Sprintf("%d", values[0])
	} else if len(values) == 2 {
		if k == 1 {
			return fmt.Sprintf("%d%d", values[0], values[1])
		} else {
			return fmt.Sprintf("%d%d", values[1], values[0])
		}
	} else {
		base := getFactorial(len(values) - 1)
		groupId := k / base
		idx := k % base
		if idx == 0 {
			groupId -= 1
			idx = base
		}
		return fmt.Sprintf("%d", values[groupId]) + getPermutationInternal(append(values[:groupId], values[groupId+1:]...), idx)
	}
}

func getFactorial(n int) int {
	return factorials[n-1]
}
