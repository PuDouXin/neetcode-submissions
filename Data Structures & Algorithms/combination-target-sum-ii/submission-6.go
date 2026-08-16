func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	current := []int{}
	
	sort.Ints(candidates)

	var backpacking func(i int, n int)
	backpacking = func(i int, n int){
		if n == target{
			tmp := make([]int, len(current))
			copy(tmp, current)
			res = append(res,tmp)
			return
		}
		if n > target{
			return
		}
		seen := make(map[int]bool)
		for s := i; s<len(candidates); s++{
			if seen[candidates[s]]{
				continue
			}
			seen[candidates[s]] = true
			current = append(current, candidates[s])
			backpacking(s+1,n+candidates[s])
			
			current = current[:len(current)-1]
			
		}

	}
	backpacking(0,0)
	return res
}

