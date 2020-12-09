package problem_321

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var mss []int
	for l := 0; l <= k; l++ {
		r := k - l
		if l <= len(nums1) && r <= len(nums2) {
			ls := maxSubSeq(nums1, l)
			rs := maxSubSeq(nums2, r)
			ss := merge(ls, rs)
			if mss == nil || comp2(ss, mss) {
				mss = ss
			}
		}
	}
	return mss
}

func comp2(nums1, nums2 []int) bool {
	for i, v := range nums1 {
		r := v - nums2[i]
		if r != 0 {
			return r > 0
		}
	}
	return false
}

func merge(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	ss := make([]int, 0, len(nums1)+len(nums2))
	lpos, rpos := 0, 0
	for lpos < len(nums1) || rpos < len(nums2) {
		if comp(nums1, nums2, lpos, rpos) {
			ss = append(ss, nums1[lpos])
			lpos++
		} else {
			ss = append(ss, nums2[rpos])
			rpos++
		}
	}
	return ss
}

func comp(nums1, nums2 []int, lpos, rpos int) bool {
	for lpos < len(nums1) && rpos < len(nums2) {
		r := nums1[lpos] - nums2[rpos]
		if r != 0 {
			return r > 0
		}
		lpos++
		rpos++
	}
	return len(nums1)-lpos > len(nums2)-rpos
}

func maxSubSeq(nums []int, k int) []int {
	if k > 0 {
		stack := make([]int, 0, k)
		dropCount := len(nums) - k
		for _, v := range nums {
			for len(stack) > 0 && stack[len(stack)-1] < v && dropCount > 0 {
				stack = stack[:len(stack)-1]
				dropCount--
			}
			if len(stack) < k {
				stack = append(stack, v)
			} else {
				dropCount--
			}
		}
		return stack
	}
	return ([]int)(nil)
}
