func maxProfit(prices []int) int {
    n := len(prices)
	memo := make(map[[2]int]int,n)
	
	var dp func(int,bool)int
	//true: allow to buy; false:must sell
	dp = func(i int,buying bool)int{
		if i>=n{
			return 0
		}
		key := [2]int{i,boolToInt(buying)}
		if v,e := memo[key]; e{
			return v
		}
		cooldown := dp(i+1,buying)
		if buying{
			memo[key]= max(cooldown, dp(i+1,false) - prices[i])
		}else{
			memo[key]= max(cooldown, dp(i+2,true)+prices[i])
		}
		return memo[key]
	}
	return dp(0,true)
}
func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}