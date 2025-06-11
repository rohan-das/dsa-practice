package main

func containsNearbyDuplicate(nums []int, k int) bool {
	indices := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if idx, ok := indices[nums[i]]; ok {
			if i-idx <= k {
				return true
			}
		}
		indices[nums[i]] = i
	}
	return false
}
