func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
	for i := 0; i< numCourses; i++{
		preMap[i] = []int{}
	}
	for _, prereq := range prerequisites{
		cur, pre := prereq[0], prereq[1]
		preMap[cur] = append(preMap[cur], pre)
	}
	visiting := make(map[int]bool)

	var dfs func(int)bool
	dfs = func(cur int)bool{
		if visiting[cur]{
			return false
		}
		if len(preMap[cur]) ==0{
			return true
		}
		visiting[cur] = true

		for _, pre := range preMap[cur]{
			if !dfs(pre){
				return false
			}
		}
		visiting[cur]= false
		preMap[cur]=[]int{}
		return true
	}

	for i := 0; i< numCourses; i++{
		if !dfs(i){
			return false
		}
	}
	return true
}
