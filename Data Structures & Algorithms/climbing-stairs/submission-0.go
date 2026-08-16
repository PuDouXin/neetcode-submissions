func climbStairs(n int) int {
    cache := make([]int,n)
	for i := range cache{
		cache[i] = -1
	}
	var dfs func(int)int
	dfs = func(i int) int{
		if i ==n{
			return 1
		}
		if i > n {
			return 0
		}
		if cache[i] != -1{
			return cache[i]
		}
		cache[i] = dfs(i+1) + dfs(i+2)
		return cache[i]
	}
	return dfs(0)
	
}
