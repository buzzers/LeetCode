package problem_40

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	// 排序，用于后续剪枝
	sort.Ints(candidates)
	steps := buildSteps(candidates)
	return combinationSum2Internal(candidates, target, steps)
}

func combinationSum2Internal(candidates []int, target int, steps []int) [][]int {
	var results [][]int
	for i := 0; i < len(candidates); {
		if candidates[i] > target {
			// 如果当前候选数字过大则结束循环
			break
		} else if candidates[i] == target {
			// 如果相等则添加为结果并结束循环
			results = append(results, []int{candidates[i]})
			break
		} else if i+1 < len(candidates) && target-candidates[i] >= candidates[i+1] {
			// 如果目标数字减去候选数字大于等于最小的可用候选数字
			// 则递归搜索结果
			for _, result := range combinationSum2Internal(candidates[i+1:], target-candidates[i], steps[i+1:]) {
				results = append(results, append(result, candidates[i]))
			}
		}
		// 选择下一个候选数字
		i += steps[i]
	}
	return results
}

// 构建下一个候选数字步进表
func buildSteps(candidates []int) []int {
	steps := make([]int, len(candidates))
	pos := 0
	for i := 1; i < len(steps); i++ {
		if candidates[i] != candidates[pos] {
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
