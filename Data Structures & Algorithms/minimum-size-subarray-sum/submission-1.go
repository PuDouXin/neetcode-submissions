func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0{
		return 0
	}
	res := math.MaxInt
	i, j, sum := 0,0,0
	
	for j<= len(nums)-1{
		sum += nums[j]
		for sum >=target{
			if j-i+1 < res{
				res = j-i+1
			}
			sum -= nums[i]
			i++
		
		}
	
		j++
	}
	if res == math.MaxInt{
		return 0
	}
	return res
}
