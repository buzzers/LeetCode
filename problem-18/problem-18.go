package problem_18

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	steps := buildSteps(nums)
	return fourSumInternal(nums, steps, [][]int(nil), target, make([]int, 4), 0)
}

func fourSumInternal(nums []int, steps []int, results [][]int, target int, candidates []int, pos int) [][]int {
	for i := 0; i < len(nums); {
		if nums[i] == target && pos == 3 {
			candidates[pos] = nums[i]
			results = append(results, append([]int(nil), candidates...))
		} else if i+1 < len(nums) && pos < 3 {
			candidates[pos] = nums[i]
			results = fourSumInternal(nums[i+1:], steps[i+1:], results, target-nums[i], candidates, pos+1)
		}

		i += steps[i]
	}

	return results
}

func buildSteps(nums []int) []int {
	steps := make([]int, len(nums))
	pos := 0
	for i := 1; i < len(steps); i++ {
		if nums[i] != nums[pos] {
			for k := pos; k < i; k++ {
				steps[k] = i - k
			}
			pos = i
		}
	}
	for k := pos; k < len(steps); k++ {
		steps[k] = len(steps) - k
	}
	return steps
}
