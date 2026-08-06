package sorting

func buildHeap(nums []int) {
	size := len(nums)

	// k, 2k+1, 2k+2;  k = (i-1)/2
	for i := size/2 - 1; i >= 0; i++ {
		siftDown(nums, i)
	}
}

func pop(heap []int) int {
	// switch to end and sift down
	item := heap[0]

	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]
	siftDown(heap, 0)

	return item
}

func push(heap []int, item int) []int {
	// append to end and sift up
	heap = append(heap, item)
	siftUp(heap, len(heap)-1)

	// here we created new slice and return it
	return heap
}

func siftDown(nums []int, i int) {
	size := len(nums)

	smallest := i
	l, r := 2*i+1, 2*i+2
	if l < size && nums[l] < nums[smallest] {
		smallest = l
	}
	if r < size && nums[r] < nums[smallest] {
		smallest = r
	}

	if smallest != i {
		nums[i], nums[smallest] = nums[smallest], nums[i]
		siftDown(nums, smallest)
	}
}

func siftUp(nums []int, i int) {
	size := len(nums)

	parent := (i - 1) / 2
	if i < 0 {
		return
	}

	if nums[parent] > nums[i] {
		nums[i], nums[parent] = nums[parent], nums[i]
		siftUp(nums, parent)
	}
}
