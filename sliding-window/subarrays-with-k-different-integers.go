package main

func subarraysWithKDistinct(nums []int, k int) int {
	return atMostKDistinctIntSubArr(nums, k) - atMostKDistinctIntSubArr(nums, k-1)
}

func atMostKDistinctIntSubArr(nums []int, k int) int {
	left, subArr := 0, 0
	freq := make(map[int]int)

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		for len(freq) > k {
			if freq[nums[left]] == 1 {
				delete(freq, nums[left])
			} else {
				freq[nums[left]]--
			}
			left++
		}
		subArr += right - left + 1
	}

	return subArr
}
