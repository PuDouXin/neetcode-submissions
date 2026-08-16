func combinationSum(nums []int, target int) [][]int {
	if len(nums) == 0{
		return [][]int{}
	}
    var res [][]int
	current := []int{}
	var backpacking func(i int, n int)
	backpacking = func(i int, n int){
		
		if n == target{
			
			tmp := make([]int, len(current))
			copy(tmp, current)
			res = append(res, tmp)
			return
		}
		if n > target || i >=len(nums){return}
		
			current = append(current,nums[i])
			backpacking(i, n+nums[i])
			current = current[:len(current)-1]
			backpacking(i+1, n)


		
	}

	backpacking(0,0)
	return res
}
