func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := range dp {
		dp[i][0] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			lastMinArg := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				lastMinArg++
			}
			dp[i][j] = min(dp[i][j], lastMinArg)
		}
	}
	return dp[len(word1)][len(word2)]

}
