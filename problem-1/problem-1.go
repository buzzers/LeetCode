package problem_1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		x := target - n
		if a, ok := m[x]; ok {
			return []int{a, i}
		} else {
			m[n] = i
		}
	}
	return nil
}
