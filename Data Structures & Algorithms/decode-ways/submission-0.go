func numDecodings(s string) int {
    memo := make([]int, len(s))
	for i := range s{
		memo[i] = -1
	}

	var dp func(i int)int
	dp = func(i int)int{
		if i==len(s){
			return 1
		}
		if s[i] == '0'{
			return 0
		}
		
		if memo[i]!=-1{
			return memo[i]
		}

		res := dp(i+1)
		if i+1<len(s) &&(s[i] =='1' || (s[i]=='2' && s[i+1]<='6')){
			res+= dp(i+2)
		}
		memo[i] = res
		return res

	}

	return dp(0)
}
