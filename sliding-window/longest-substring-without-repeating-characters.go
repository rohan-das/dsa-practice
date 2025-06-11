package main

func lengthOfLongestSubstring(s string) int {
	left, maxLen := 0, 0
	freq := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for len(freq) != right-left+1 {
			if freq[s[left]] == 1 {
				delete(freq, s[left])
			} else {
				freq[s[left]]--
			}
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
