package main

import "math"

func minSwaps(data []int) int {
	maxOnes := 0
	for i := 0; i < len(data); i++ {
		maxOnes += data[i]
	}

	if maxOnes == 0 || maxOnes == len(data) {
		return 0
	}

	ones, left, minSwap := 0, 0, math.MaxInt
	for right := 0; right < len(data); right++ {
		if data[right] == 1 {
			ones++
		}

		// maxOnes here is the maximum window
		if right-left+1 == maxOnes {
			minSwap = min(minSwap, maxOnes-ones)
			if data[left] == 1 {
				ones--
			}
			left++
		}

	}

	return minSwap
}
