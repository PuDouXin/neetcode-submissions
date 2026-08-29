func findTargetSumWays(nums []int, target int) int {
    memo := make(map[[2]int]int)
	var dp func(int,int)int

	dp = func(i, total int)int{
		if i==len(nums) {
			if total == target{
				return 1
			}
			return 0
		}
		key := [2]int{i,total}
		if v,exist := memo[key]; exist {
			return v
		}
		memo[key] = dp(i+1, total+nums[i])+dp(i+1,total-nums[i])
		return memo[key]
		
	}
	return dp(0,0)
}
