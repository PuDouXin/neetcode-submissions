func change(amount int, coins []int) int {
	sort.Ints(coins)
    memo := make([][]int, len(coins)+1)
	for i := range memo{
		memo[i] = make([]int, amount+1)
		for j := range memo[i]{
			memo[i][j]=-1
		}
	}

	var dp func(int,int)int
	dp = func(i, a int)int{
		if a ==0{
			return 1
		}
		if i>=len(coins){
			return 0
		}
		if memo[i][a]!=-1{
			return memo[i][a]
		}
		res :=0
		if a>=coins[i]{
			res = dp(i+1,a)+dp(i, a-coins[i])
		}
		memo[i][a] = res
		return res

	}
	return dp(0,amount)
}
