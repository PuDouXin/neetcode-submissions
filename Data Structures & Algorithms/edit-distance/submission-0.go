func minDistance(word1 string, word2 string) int {
    m := len(word1)
	n := len(word2)
	memo := make([][]int,m+1)
	for i := range memo{
		memo[i] = make([]int,n+1)
		for j := range memo[i]{
			memo[i][j]=-1
		}
	}
	var dp func(int,int)int
	dp = func(i,j int)int{
		if i == m{
			return n-j
		}
		if j == n{
			return m-i
		}
		if memo[i][j]!=-1{
			return memo[i][j]
		}

		if word1[i]==word2[j]{
			memo[i][j]=dp(i+1,j+1)
		}else{
			memo[i][j] = 1+min(min(dp(i+1,j),dp(i,j+1)), dp(i+1,j+1))
		}
		return memo[i][j]
	}
	return dp(0,0)
}
