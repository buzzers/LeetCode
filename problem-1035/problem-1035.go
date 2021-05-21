package problem1035

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
