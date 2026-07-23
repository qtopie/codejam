package leetcode

func reverseAndMinimum(nums []int, start, end int) (index int) {
	index = start
	for i := start + 1; i <= end; i++ {
		if nums[i] < nums[index] {
			index = i
		}
	}

	reverse(nums, start, index)
	return index
}

func reverse(nums []int, start, end int) []int {
	if start == end {
		return nums
	}

	for i := 0; i <= (end-start)/2+start; i++ {
		nums[i], nums[end-i] = nums[end-i], nums[i]
	}

	return nums
}

func resortCost(L []int) (cost int) {
	end := len(L) - 1
	for i := 1; i <= end; i++ {
		j := reverseAndMinimum(L, i-1, end) + 1
		cost += j - i + 1
	}

	return cost
}

// func main() {

// 	var n int
// 	fmt.Scanln(&n)

// 	for i := 1; i <= n; i++ {
// 		var size int
// 		fmt.Scanf("%d", &size)
// 		arr := make([]int, size)
// 		y := make([]interface{}, size)
// 		for k := range arr {
//         y[k] = &arr[k]
//     }
// 		fmt.Scanln(y...)
// 		cost := resortCost(arr)
// 		fmt.Printf("Case #%d: %d\n", i, cost)
// 	}
// }
