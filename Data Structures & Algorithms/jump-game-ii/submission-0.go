func jump(nums []int) int {
    memo := make(map[int]int)

	var dfs func(int) int
	dfs = func(i int)int{
		if val, existed := memo[i]; existed{
			return val
		}
		if i == len(nums)-1{
			return 0
		}
		if nums[i]==0{
			return 100000
		}
		res := 100000
		end := min(i+nums[i]+1,len(nums))
		for j:= i+1;j<end;j++{
			res = min(res, 1+dfs(j))
		}
		memo[i]=res
		return res 
	}
	return dfs(0)
}
