func subsets(nums []int) [][]int {
	res:=[][]int{}
	current := []int{}
	var backpacking func(start int)

	backpacking = func(start int){
		tmp := make([]int, len(current))
		copy(tmp, current)
		res = append(res, tmp)

		for i := start; i < len(nums); i++{
			current = append(current, nums[i])
			backpacking(i+1)
			current = current[:len(current)-1]
		}
	}

	backpacking(0)
	return res


}
