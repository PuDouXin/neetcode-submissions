func coinChange(coins []int, amount int) int {
	memo := make(map[int]int)
	memo[0] = 0

	var dp func(int) int
	dp = func(a int) int {
		if val, exist := memo[a]; exist {
			return val
		}
		res := math.MaxInt32
		for _, coin := range coins {
			if a-coin >= 0 {
				res = min(res, 1+dp(a-coin))
			}
		}
		memo[a] = res
		return res
	}
	res := dp(amount)
	if res == math.MaxInt32 {
		return -1
	}
	return res

}
