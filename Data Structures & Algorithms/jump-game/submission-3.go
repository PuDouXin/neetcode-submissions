func canJump(nums []int) bool {
    memo := make(map[int]bool)
	var dfs func(int)bool
	dfs = func(i int)bool{
		if i >=len(nums)-1{
			return true
		}
		if res, existed := memo[i]; existed{
			return res
		}
		if nums[i]==0{
			return false
		}
		end := min(len(nums),i+nums[i]+1)
		for j:= i+1; j<end; j++{
			if dfs(j){
				memo[i] = true
				return true
			}
		}
		memo[i] = false
		return false
	}
	return dfs(0)
}
