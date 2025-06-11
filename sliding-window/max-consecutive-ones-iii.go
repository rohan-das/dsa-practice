package main

func longestOnes(nums []int, k int) int {
	var left, zeroes, maxOnes int
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroes++
			for zeroes > k {
				if nums[left] == 0 {
					zeroes--
				}
				left++
			}
		}
		maxOnes = max(maxOnes, right-left+1)
	}

	return maxOnes
}
