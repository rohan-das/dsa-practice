package main

import "math"

// Original function name was minSwaps : changed it to MinSwaps
func MinSwaps(nums []int) int {
	maxOnes, n := 0, len(nums)
	for i := 0; i < n; i++ {
		maxOnes += nums[i]
	}

	if maxOnes == 0 || maxOnes == n {
		return 0
	}

	left, ones, minSwap := 0, 0, math.MaxInt
	for right := 0; right < 2*n; right++ {
		if nums[right%n] == 1 {
			ones++
		}

		// maxOnes here is the maximum window
		if right-left+1 == maxOnes {
			minSwap = min(minSwap, maxOnes-ones)
			if nums[left%n] == 1 {
				ones--
			}
			left++
		}
	}

	return minSwap
}
