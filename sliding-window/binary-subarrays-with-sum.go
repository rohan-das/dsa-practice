package main

func numSubarraysWithSum(nums []int, goal int) int {
	return atMostKTimesSumEqualToGoal(nums, goal) - atMostKTimesSumEqualToGoal(nums, goal-1)
}

func atMostKTimesSumEqualToGoal(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}

	var left, count, sum int
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum > goal {
			sum -= nums[left]
			left++
		}
		// Count valid subarrays ending at right
		count += right - left + 1
	}

	return count
}
