package main

func maxScore(cardPoints []int, k int) int {
	n, sum := len(cardPoints), 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	maxScore := sum
	for i := 1; i <= k; i++ {
		sum = sum + cardPoints[n-i] - cardPoints[k-i]
		maxScore = max(maxScore, sum)
	}

	return maxScore
}
