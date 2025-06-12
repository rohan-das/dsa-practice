package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	left, subArrCount, prod := 0, 0, 1
	for right := 0; right < len(nums); right++ {
		prod = prod * nums[right]
		for left <= right && prod >= k {
			prod = prod / nums[left]
			left++
		}
		subArrCount += right - left + 1
	}

	return subArrCount
}
