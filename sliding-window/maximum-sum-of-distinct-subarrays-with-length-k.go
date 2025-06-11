package main

func maximumSubarraySum(nums []int, k int) int64 {
	var sum, maxSum int64
	left, freq := 0, make(map[int]int)

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		sum += int64(nums[right])
		if right-left+1 == k {
			if len(freq) == k {
				maxSum = max(maxSum, sum)
			}
			if freq[nums[left]] == 1 {
				delete(freq, nums[left])
			} else {
				freq[nums[left]]--
			}
			sum -= int64(nums[left])
			left++
		}
	}

	return maxSum
}
