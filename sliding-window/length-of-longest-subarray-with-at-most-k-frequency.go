package main

func maxSubarrayLength(nums []int, k int) int {
	left, longestGoodSubArr := 0, 0
	freq := make(map[int]int)

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		for freq[nums[right]] > k {
			if freq[nums[left]] == 1 {
				delete(freq, nums[left])
			} else {
				freq[nums[left]]--
			}
			left++
		}
		longestGoodSubArr = max(longestGoodSubArr, right-left+1)
	}

	return longestGoodSubArr
}
