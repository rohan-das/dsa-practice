package main

func totalFruit(fruits []int) int {
	left, maxFruits := 0, 0
	freq := make(map[int]int)

	for right := 0; right < len(fruits); right++ {
		freq[fruits[right]]++
		for len(freq) > 2 {
			if freq[fruits[left]] == 1 {
				delete(freq, fruits[left])
			} else {
				freq[fruits[left]]--
			}
			left++
		}
		maxFruits = max(maxFruits, right-left+1)
	}

	return maxFruits
}
