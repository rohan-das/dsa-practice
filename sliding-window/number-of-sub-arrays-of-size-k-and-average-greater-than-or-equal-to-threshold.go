package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	var left, sum, subArr int
	for right := 0; right < len(arr); right++ {
		sum += arr[right]
		if right-left+1 == k {
			avg := sum / k
			if avg >= threshold {
				subArr++
			}
			sum -= arr[left]
			left++
		}
	}

	return subArr
}
