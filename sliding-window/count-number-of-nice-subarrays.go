package main

func numberOfSubarrays(nums []int, k int) int {
	return atMostKTimesOddNums(nums, k) - atMostKTimesOddNums(nums, k-1)
}

func atMostKTimesOddNums(nums []int, k int) int {
	left, count, odd := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right]%2 != 0 {
			odd++
		}

		// Shrink window if more than k odd numbers
		for odd > k {
			if nums[left]%2 != 0 {
				odd--
			}
			left++
		}

		// Count valid subarrays ending at right
		count += right - left + 1
	}
	return count
}
