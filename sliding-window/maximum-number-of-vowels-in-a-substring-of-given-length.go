package main

func maxVowels(s string, k int) int {
	left, vowel, maxVowel := 0, 0, 0
	for right := 0; right < len(s); right++ {
		if isVowel(s[right]) {
			vowel++
		}
		maxVowel = max(maxVowel, vowel)
		if right-left+1 == k {
			if isVowel(s[left]) {
				vowel--
			}
			left++
		}
	}

	return maxVowel
}

func isVowel(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	return false
}
