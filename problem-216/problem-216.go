package problem_216

func combinationSum3(k int, n int) [][]int {
	candidates := make([]int, 9)
	for i := range candidates {
		candidates[i] = i + 1
	}
	return combinationSum3Internal(candidates, n, k)
}

func combinationSum3Internal(candidates []int, target int, step int) [][]int {
	var results [][]int
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			// 如果当前候选数字过大则结束循环
			break
		} else if step == 1 && candidates[i] == target {
			// 如果相等且步数符合要求则添加为结果并结束循环
			results = append(results, []int{candidates[i]})
			break
		} else if step > 1 && i+1 < len(candidates) && target-candidates[i] >= candidates[i+1] {
			// 如果目标数字减去候选数字大于等于最小的可用候选数字且剩余步数足够
			// 则递归搜索结果
			for _, result := range combinationSum3Internal(candidates[i+1:], target-candidates[i], step-1) {
				results = append(results, append(result, candidates[i]))
			}
		}
	}
	return results
}
