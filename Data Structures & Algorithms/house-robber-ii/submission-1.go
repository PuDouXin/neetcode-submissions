func rob(nums []int) int {
	if len(nums) ==1{
		return nums[0]
	}
    memo := make([][2]int, len(nums))
	for i := range memo{
		memo[i][0] = -1
		memo[i][1] = -1
	}

	var dfs func(i int, flag int) int
	dfs = func(i int, flag int)int{
		if i>=len(nums) || (flag==1 && i ==len(nums)-1){
			return 0
		}
		if memo[i][flag]!=-1{
			return memo[i][flag]
		}
		tmp := flag
		if i==0{
			tmp =1
		}
		memo[i][flag]= max(nums[i]+dfs(i+2,tmp), dfs(i+1,flag))
		return memo[i][flag]
	}
	return max(dfs(0, 1), dfs(1, 0))
}
