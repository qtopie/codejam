package dynamic_programming

// 有两个序列 a,b，大小都为 n,序列元素的值任意整数，无序； 要求：通过交换 a,b 中的元素，使[序列 a 元素的和]与[序列 b 元素的和]之间的差最小

// 可以转化成从2n个数里选n个数,使得这n个数的和最接近sum/2 (即全部和的一半) 可以只考虑 <= sum/2的情况
// 转换成01背包问题: 背包容量为sum/2, 每件物品的价值为w, 重量也为w,  但是需要选n个物品(背包解的特殊子集), 使其总重量接近sum/2
// f(i,j,k) = max(f(i-1, j, k), f(i-1, j-1, k - w[i]) + w[i])
// i个物品,选j个,重量为k, f求其最大价值
func minimizeDifference(nums1, nums2 []int) int {
	n := len(nums1)
	sum := 0
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i]
	}

	for i := 0; i < len(nums2); i++ {
		sum += nums2[i]
	}

	capacity := (sum + 1) / 2

	dp := make([][]int, n+1)
	for j := 0; j <= 2*n; j++ {
		dp = append(dp, make([]int, capacity+1))
	}

	for i := 1; i <= 2*n; i++ {
		j := i
		var w int
		if j > n {
			j = n
			w = nums2[i-n-1]
		} else {
			w = nums1[i-1]
		}

		for ; j >= 1; j-- {
			for k := capacity; k >= 1; k-- {
				if k > w && dp[j][k] < dp[j-1][k-w]+w {
					dp[j][k] = dp[j-1][k-w] + w
				}
			}
		}
	}

	return sum - 2*dp[n][capacity]
}

func _brute_force_difference(nums1, nums2 []int) int {
	n := len(nums1)
	sum := 0
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i]
	}

	for i := 0; i < len(nums2); i++ {
		sum += nums2[i]
	}

	half := (sum + 1) / 2

	// calculate combinations
	size := len(nums1) * 2
	combinations := combine(size, size/2)

	// for each combinations caculate result and get minimal one
	minVal := sum
	for _, cmb := range combinations {
		val := 0
		for _, k := range cmb {
			if k <= n {
				val += nums1[k-1]
			} else {
				val += nums2[k-n-1]
			}

			if val > half {
				break
			}
		}

		if val < half && val < minVal {
			minVal = val
		}
	}

	return minVal
}

func combine(n int, k int) [][]int {
	// initialize first line [] [][]
	current := make([][][]int, k)
	current[0] = [][]int{{1}}

	for i := 2; i <= n; i++ {
		j := k
		// dp
		for ; j > 1; j-- {
			if j == i {
				// calculate C_i^j, append j
				c := make([]int, j-1)
				copy(c, current[j-2][0]) // copy C_j-1^j-1, only one combination
				c = append(c, j)
				current[j-1] = [][]int{c}
				continue
			}

			cmb := current[j-1]
			for _, c := range current[j-2] {
				newC := make([]int, len(c))
				copy(newC, c)
				newC = append(newC, i)
				cmb = append(cmb, newC)
			}
			current[j-1] = cmb
		}

		// j == 1, add i to result
		current[j-1] = append(current[j-1], []int{i})
	}

	return current[k-1]
}
