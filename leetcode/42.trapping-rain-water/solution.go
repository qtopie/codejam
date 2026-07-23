package leetcode

// 相似算法题  Container With Most Water https://leetcode.com/problems/container-with-most-water/description/

// trap = min(lmax, rmax) - height[i]
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	sum := 0
	lMax, rMax := 0, 0
	for l, r := 0, len(height)-1; l <= r; {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])

		if lMax < rMax {
			sum += lMax - height[l]
			l++
		} else {
			sum += rMax - height[r]
			r--
		}
	}

	return sum
}

type PriorityQueue struct {
	val [][3]int
}

func (pq PriorityQueue) Len() int {
	return len(pq.val)
}

func (pq *PriorityQueue) Init(val [][3]int) {
	pq.val = val
	size := pq.Len()

	for i := size/2 - 1; i >= 0; i-- {
		pq.siftDown(i)
	}
}

func (pq *PriorityQueue) Push(item [3]int) {
	pq.val = append(pq.val, item)
	pq.siftUp(pq.Len() - 1)
}

func (pq *PriorityQueue) Pop() [3]int {
	item := pq.val[0]

	size := pq.Len()
	pq.val[0] = pq.val[size-1]
	pq.val = pq.val[:size-1]

	pq.siftDown(0)
	return item
}

func (pq *PriorityQueue) siftUp(i int) {
	parent := (i - 1) / 2
	if parent >= 0 && pq.val[i][2] < pq.val[parent][2] {
		pq.val[i], pq.val[parent] = pq.val[parent], pq.val[i]
		pq.siftUp(parent)
	}
}

func (pq *PriorityQueue) siftDown(i int) {
	size := pq.Len()
	smallest := i
	l, r := 2*i+1, 2*i+2

	if l < size && pq.val[l][2] < pq.val[smallest][2] {
		smallest = l
	}

	if r < size && pq.val[r][2] < pq.val[smallest][2] {
		smallest = r
	}

	if smallest != i {
		pq.val[smallest], pq.val[i] = pq.val[i], pq.val[smallest]

		pq.siftDown(smallest)
	}
}

func trapRainWater(heightMap [][]int) int {
	sum := 0
	x, y := len(heightMap), len(heightMap[0])
	if x < 3 || y < 3 {
		return sum
	}

	visited := make([][]bool, x)
	for i := 0; i < x; i++ {
		visited[i] = make([]bool, y)
	}

	var pq PriorityQueue
	val := make([][3]int, 0)
	for i := 0; i < x; i++ {
		val = append(val, [3]int{i, 0, heightMap[i][0]})
		val = append(val, [3]int{i, y - 1, heightMap[i][y-1]})
		visited[i][0] = true
		visited[i][y-1] = true
	}

	for j := 1; j < y-1; j++ {
		val = append(val, [3]int{0, j, heightMap[0][j]})
		val = append(val, [3]int{x - 1, j, heightMap[x-1][j]})
		visited[0][j] = true
		visited[x-1][j] = true
	}

	pq.Init(val)

	biases := []int{-1, 0, 1, 0, -1}

	for pq.Len() > 0 {
		// min height
		item := pq.Pop()

		for i := 0; i < 4; i++ {
			dx, dy := item[0]+biases[i], item[1]+biases[i+1]
			if dx >= 0 && dx < x && dy >= 0 && dy < y && !visited[dx][dy] {
				maxH := item[2]

				if item[2] > heightMap[dx][dy] {
					sum += item[2] - heightMap[dx][dy]
				} else {
					maxH = heightMap[dx][dy]
				}

				pq.Push([3]int{dx, dy, maxH})
				visited[dx][dy] = true
			}
		}
	}

	return sum
}
