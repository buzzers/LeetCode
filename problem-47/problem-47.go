package problem_47

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return permuteUniqueInternal(nums, make([]bool, len(nums)), make([]int, 0, len(nums)), nil)
}

func permuteUniqueInternal(nums []int, visited []bool, values []int, results [][]int) [][]int {
	if len(nums) == len(values) {
		results = append(results, append([]int(nil), values...))
		return results
	}
	for i, num := range nums {
		if visited[i] || i > 0 && !visited[i-1] && num == nums[i-1] {
			continue
		}
		values = append(values, num)
		visited[i] = true
		results = permuteUniqueInternal(nums, visited, values, results)
		visited[i] = false
		values = values[:len(values)-1]
	}
	return results
}
