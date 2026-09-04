func maxSubArray(nums []int) int {
    res, curSum := nums[0],0
	for _, num := range nums{
		if curSum < 0{
			curSum = 0
		}
		curSum +=num
		res = max(res, curSum)
	}
	return res
}

