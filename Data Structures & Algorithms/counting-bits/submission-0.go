func countBits(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0
	offset := 1
	for i:=1;i<=n;i++{
		if i==2*offset{
			offset = i
		}
		dp[i] = 1+dp[i-offset]
	}
	return dp
}
