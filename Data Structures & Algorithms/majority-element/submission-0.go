func majorityElement(nums []int) int {
    countS :=make(map[int]int)
	for _,n:= range nums{
		countS[n]++
	}
	
	for k,v:=range countS{
		if v>len(nums)/2{
			return k
		}
	}
	return 0
}
