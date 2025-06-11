package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	left, longestSubStr := 0, 0
	freq := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for len(freq) > k {
			if freq[s[left]] == 1 {
				delete(freq, s[left])
			} else {
				freq[s[left]]--
			}
			left++
		}
		longestSubStr = max(longestSubStr, right-left+1)
	}

	return longestSubStr
}
