package main

func countGoodSubstrings(s string) int {
	left, count := 0, 0
	freq := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		if right-left+1 == 3 {
			if len(freq) == 3 {
				count++
			}
			if freq[s[left]] == 1 {
				delete(freq, s[left])
			} else {
				freq[s[left]]--
			}
			left++
		}
	}

	return count
}
