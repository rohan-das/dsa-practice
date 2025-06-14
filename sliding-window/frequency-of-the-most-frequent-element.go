package main

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left, maxFreq := 0, 1

	for right := 1; right < len(nums); right++ {
		k -= (nums[right] - nums[right-1]) * (right - left)
		for k < 0 {
			k += nums[right] - nums[left]
			left++
		}
		maxFreq = max(maxFreq, right-left+1)
	}

	return maxFreq
}
