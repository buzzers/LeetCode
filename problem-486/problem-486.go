package problem_486

func PredictTheWinner(nums []int) bool {
	// states[i][j] 代表剩余数字是从 i 到 j 时当前的玩家的分数差值

	states := make([][]int, len(nums))
	for i := range states {
		states[i] = make([]int, len(nums))
	}
	for i := range nums {
		// 仅剩一个数分数差值便是剩余的数
		states[i][i] = nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			// 状态转移方程
			// 从 i 到 j 的差值 = max(左边的数值减去从 i+1 到 j 的差值, 右边的数值减去从 i 到 j-1 的差值)
			states[i][j] = max(nums[i]-states[i+1][j], nums[j]-states[i][j-1])
		}
	}
	return states[0][len(nums)-1] >= 0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
