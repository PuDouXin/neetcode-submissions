func maxProfit(prices []int) int {
	res := 0
	for i:=0;i<len(prices);i++{
		buy := prices[i]
		for j:= i+1; j<len(prices);j++{
			sell := prices[j]
			if (sell-buy)>res{
				res = sell-buy
			}
		}
		
	}
	return res
}
