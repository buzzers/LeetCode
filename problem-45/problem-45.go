package problem_45

func jump(nums []int) int {
	step := 0
	// 从尾部开始扫描
	// 寻找能一步跳到 si 的位置，此位置便是 i 到 si 的最优解
	si := len(nums) - 1
	for si > 0 {
		l := si - 1
		for i := si - 1; i >= 0; i-- {
			if nums[i] >= si-i {
				l = i
			}
		}
		si = l
		step++
	}
	return step
}
