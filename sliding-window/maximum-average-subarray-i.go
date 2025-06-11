package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	maxAvg := math.Inf(-1)
	sum, left := 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		if right-left+1 == k {
			maxAvg = math.Max(maxAvg, float64(sum)/float64(k))
			sum -= nums[left]
			left++
		}
	}

	return maxAvg
}
