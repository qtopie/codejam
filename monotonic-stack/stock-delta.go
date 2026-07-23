package main

import "fmt"

// 5 minutes < 5%
func isSteady(prices []float64) bool {
	if len(prices) <= 1 {
		return true
	}

	windowSize := min(len(prices), 300)
	// descending maxDeque, the newest larger element is preferred
	maxDeque := []int{0}
	// asending minDequeu, the newest smaller element is preferred
	minDequeu := []int{0}

	// fill queues
	for i := 1; i <= len(prices); i++ {
		// new element into queue
		for prices[i] >= maxDeque[len(maxDeque)-1] {
			// pop right
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, i)

		for prices[i] <= minDeque[len(minDeque)-1] {
			// pop right
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, i)

		if i >= windowSize {
			// check window
			if !checkDeltaMoreThanFivePercents(minDeque[0], maxDeque[0]) {
				return false
			}

			// drop oldest one
			if i-maxDeque[0] > windowSize {
				maxDeque = maxDeque[1:]
			}
			if i-minDeque[0] > windowSize {
				minDeque = minDeque[1:]
			}
		}
	}

	return true
}

func checkDeltaMoreThanFivePercents(minVal, maxVal float64) bool {
	// fix me
	return (maxVal-minVal)/minVal > 0.5
}

func main() {
	prices := make([]float64, 1000)
	for i := 0; i < len(prices); i++ {
		prices[i] = 1.0
	}

	prices[200] = 2.0

	fmt.Println(isSteady(prices))
}
