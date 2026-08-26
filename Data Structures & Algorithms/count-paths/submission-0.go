func uniquePaths(m int, n int) int {
    memo := make([][]int, m)
	for i:= range m{
		memo[i] = make([]int,n)
		for j := range n{
			memo[i][j]=-1
		}
	}

	var dp func(int,int)int
	dp = func(i,j int)int{
		if i<0 || i <0 || i>=m || j>=n{
			return 0
		}
		if i==m-1 && j==n-1{
			return 1
		}
		if memo[i][j]!=-1{
			return memo[i][j]
		}
		memo[i][j] = dp(i+1,j)+dp(i,j+1)
		return memo[i][j]
	}
	return dp(0,0)
}
