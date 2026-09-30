func missingNumber(nums []int) int {
	n := len(nums)
	res := n
	for i:= n-1;i >=0;i--{
		res ^= i^nums[i]
	}
	return res
}
