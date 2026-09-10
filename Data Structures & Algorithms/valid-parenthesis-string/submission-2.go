func checkValidString(s string) bool {
	memo := make([][]int, len(s)+1)
	for i := range memo{
		memo[i]=make([]int, len(s)+1)
		for j := range memo[i]{
			memo[i][j]=-1
		}
	}
    var dfs func(i,open int)bool
	dfs = func(i, open int)bool{
		if open <0 {
			return false
		}
		if i == len(s){
			return open ==0
		}
		if memo[i][open]!=-1{
			return memo[i][open] ==1
		}
		result := false
		if s[i] =='('{
			result = dfs(i+1,open+1)
		}else if s[i] ==')'{
			result = dfs(i+1, open-1)
		}else{
			result = (dfs(i+1,open) || dfs(i+1, open-1)|| dfs(i+1,open+1))
		}
		memo[i][open] = 1
		if !result{
			memo[i][open]=0
		}
		return result
	}

	return dfs(0,0)
}
