func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)
    memo := make([][]int,m+1)
	for i := range memo{
		memo[i] = make([]int,n+1)
		for j := range memo[i]{
			memo[i][j] = -1
		}
	}

	var dp func(int, int)int
	dp = func(i,j int)int{
		if j == n{
			return 1
		}
		if i == m{
			return 0
		}
		if memo[i][j]!=-1{
			return memo[i][j]
		}
		res := dp(i+1,j)
		if s[i] == t[j]{
			res+=dp(i+1,j+1)
		}
		memo[i][j]=res
		return res
	}
	return dp(0,0)

}
