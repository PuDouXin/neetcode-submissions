func max(a int, b int) int{
	if a>b{
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	res := 0
	l := 0
	r := 1
	for r<len(prices){
		if prices[r]>prices[l]{
			res = max(res,prices[r]-prices[l])
		}else{
			l=r
		}
		r++
	}
	return res
}
