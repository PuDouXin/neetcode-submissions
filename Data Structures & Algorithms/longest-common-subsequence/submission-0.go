func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	memo := make([][]int,m)
	for i := range m{
		memo[i] = make([]int, n)
	}

	var dp func(int, int)int
	dp = func(i int,j int)int{
		if i>=m || j>=n{
			return 0
		}
		if memo[i][j]!=0{
			return memo[i][j]
		}
		if text1[i] ==text2[j]{
			memo[i][j] =1+dp(i+1,j+1)
			
		}else{
			memo[i][j] = max(dp(i,j+1),dp(i+1,j))
		}
		return memo[i][j]
		

	}
	return dp(0,0)
}
