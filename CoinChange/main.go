func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	coinsLen := len(coins)
	for i := 1; i < amount+1; i += 1 {
		dp[i] = -1
	}
	for i := 1; i < amount+1; i += 1 {
		minCoins := math.MaxInt
		found := false
		for j := 0; j < coinsLen; j += 1 {
			prev := i - coins[j]
			if prev < 0 || dp[prev] < 0 {
				continue
			}
			found = true
			minCoins = min(minCoins, dp[prev]+1)
		}
		if !found {
			dp[i] = -1
			continue
		}
		dp[i] = minCoins
	}
	return dp[amount]
}
