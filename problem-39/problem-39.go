package problem_39

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	//  排序，用于后续剪枝
	sort.Ints(candidates)
	return combinationSumInternal(candidates, target)
}

func combinationSumInternal(candidates []int, target int) [][]int {
	results := make([][]int, 0, 128)
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			// 如果当前候选数字过大，那么后续的都过大，直接结束
			break
		} else if candidates[i] == target {
			// 如果刚好则直接添加为结果
			results = append(results, []int{candidates[i]})
		} else if target-candidates[i] >= candidates[i] {
			// 如果剩余数值大于等于最小的可用候选数
			// 则递归搜索结果
			// 前面用过的候选数不再使用
			for _, result := range combinationSumInternal(candidates[i:], target-candidates[i]) {
				results = append(results, append(result, candidates[i]))
			}
		}
	}
	return results
}
