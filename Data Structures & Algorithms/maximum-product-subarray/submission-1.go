func maxProduct(nums []int) int {
    res := nums[0]

	curMin, curMax := 1,1
	for _,num := range nums{
		tmp := curMax *num
		curMax = max(tmp, max(num,num*curMin))
		curMin = min(tmp,min(num, num*curMin))
		res = max(res, curMax)
	}
	return res

}
