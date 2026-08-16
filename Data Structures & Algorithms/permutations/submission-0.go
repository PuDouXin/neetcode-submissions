func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	current := []int{}

	var backpacking func()

	backpacking = func(){
		if len(current) == len(nums){
			tmp := make([]int, len(current))
			copy(tmp, current)
			res = append(res, tmp)
			return
		}
		
		for s := 0; s< len(nums); s++{
			if used[s]{
				continue
			}
			current = append(current, nums[s])
			used[s]=true
			backpacking()
			current = current[:len(current)-1]
			used[s]=false
		}

	}
	backpacking()
	return res
}
