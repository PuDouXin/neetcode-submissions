func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}
	current := []int{}
	

	var backpacking func(i int)
	backpacking = func(i int){

		tmp := make([]int, len(current))
		copy(tmp, current)
		res = append(res, tmp)

		for s := i; s < len(nums); s++ {
			if s > i && nums[s] == nums[s-1]{
				continue
			}
			current = append(current, nums[s])
			
			backpacking(s+1)
			current = current[:len(current)-1]
			
		}
	}
	backpacking(0)
	return res
}
