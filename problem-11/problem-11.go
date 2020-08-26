package problem_11

func maxArea(height []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if len(height) == 0 {
		return 0
	}

	// ml 和 mr 分别代表当前记录的最大容积的索引
	ml, mr := 0, len(height)-1
	l, r := 0, len(height)-1
	for l < r {
		// 如果当前的容积大于记录的最大容积，则更新记录
		if min(height[l], height[r])*(r-l) > min(height[ml], height[mr])*(mr-ml) {
			ml, mr = l, r
		}
		// 向内移动短板
		// 因为容积计算公式是 min(hl, hr) * (r - l)，即容积由短板和长度决定。
		// 也就是说，如果把长板那一侧的索引向内移动，r - l 的值一定缩小，且 min(hl, hr) 的值不可能变大 (因为短板长度不变)。
		// 因此移动长板不可能获得比当前的容积更大的解。
		// 向内移动短板，则 min(hl, hr) 的值可能变大，有可能得到比当前的容器更大的解。
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return min(height[ml], height[mr]) * (mr - ml)
}
