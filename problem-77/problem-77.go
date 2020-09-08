package problem_77

func combine(n int, k int) [][]int {
	// 初始化前 k 个为 1 表示初始状态为选择前 k 个数字
	bitset := make([]bool, n)
	for i := 0; i < k; i++ {
		bitset[i] = true
	}
	cnk := c(n, k)
	results := make([][]int, cnk)
	for i := 0; i < cnk; i++ {
		results[i] = build(bitset, k)
		next(bitset)
	}
	return results
}

func next(bitset []bool) {
	// 从左向右寻找数组中的 10，并将其改为 01
	// 之后将左边的 1 移动到最左边

	counter := 0
	for pos := 0; pos < len(bitset); pos++ {
		if bitset[pos] && pos+1 < len(bitset) && !bitset[pos+1] {
			bitset[pos], bitset[pos+1] = false, true
			for k := 0; k < pos; k++ {
				if counter > 0 {
					bitset[k] = true
					counter--
				} else {
					bitset[k] = false
				}
			}
			break
		} else if bitset[pos] {
			counter++
		}
	}
}

func build(bitset []bool, k int) []int {
	result := make([]int, 0, k)
	for i, b := range bitset {
		if b {
			result = append(result, i+1)
		}
	}
	return result
}

func c(n int, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func factorial(n int) int {
	v := 1
	for i := n; i > 1; i-- {
		v *= i
	}
	return v
}
