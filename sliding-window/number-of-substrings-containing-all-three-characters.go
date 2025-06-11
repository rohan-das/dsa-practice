package main

func numberOfSubstrings(s string) int {
	n, left, subStrs := len(s), 0, 0
	freq := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for len(freq) == 3 {
			subStrs += n - right
			if freq[s[left]] == 1 {
				delete(freq, s[left])
			} else {
				freq[s[left]]--
			}
			left++
		}
	}

	return subStrs
}
