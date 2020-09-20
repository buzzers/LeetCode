package problem_78

func subsets(nums []int) [][]int {
	return subsetsInternal(nums, make([]int, 0, len(nums)), [][]int{{}})
}

func subsetsInternal(nums []int, result []int, results [][]int) [][]int {
	for i := range nums {
		result = append(result, nums[i])
		results = append(results, append([]int(nil), result...))
		results = subsetsInternal(nums[i+1:], result, results)
		result = result[:len(result)-1]
	}
	return results
}
