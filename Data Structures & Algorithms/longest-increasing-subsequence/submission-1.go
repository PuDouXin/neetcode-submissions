func lengthOfLIS(nums []int) int {
    memo := make([]int, len(nums))
	res :=1
	for i:= range memo{
		memo[i] = 1
	}
	for i := 0;i<len(nums);i++{
		for j := 0;j<i;j++{
			if nums[j]<nums[i]{
				memo[i] = max(memo[i], 1+memo[j])
				if memo[i]>res{
					res = memo[i]
				}
			}
		}
	}

	return res
}
