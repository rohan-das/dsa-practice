package main

func maximumUniqueSubarray(nums []int) int {
	freq := make(map[int]int)
	var left, sum, maxScore int

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		sum += nums[right]
		for len(freq) != right-left+1 {
			sum -= nums[left]
			if freq[nums[left]] == 1 {
				delete(freq, nums[left])
			} else {
				freq[nums[left]]--
			}
			left++
		}
		maxScore = max(maxScore, sum)
	}

	return maxScore
}
