package problem_42

func trap(height []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if len(height) <= 1 {
		return 0
	}

	// 记录总容积
	sum := 0
	// 记录最高点索引
	mi := 0
	// 从左开始扫描
	si, s := 0, 0
	for i := 1; i < len(height); i++ {
		if height[i] < height[si] {
			// 如果高度比记录的高度低，则这是一个低地，将其记录为需要减去的数值
			s += height[i]
		} else {
			// 如果高度比记录的高或相等，则这是一个容器，记录其容积
			// 容积计算公式为 min(左挡板高度, 右挡板高度) * (右挡板索引 - 左挡板索引 - 1) - 低地的容积
			sum += min(height[si], height[i])*(i-si-1) - s
			// 重置开始索引和低地的数值
			si = i
			s = 0
		}
		if height[i] >= height[mi] {
			// 更新最高点索引
			mi = i
		}
	}
	// 从右开始扫描
	si, s = len(height)-1, 0
	// 在最高点终止，因为最高点的左边已经扫描过
	for i := len(height) - 2; i >= mi; i-- {
		if height[i] < height[si] {
			s += height[i]
		} else {
			sum += min(height[si], height[i])*(si-i-1) - s
			si = i
			s = 0
		}
	}
	return sum
}
