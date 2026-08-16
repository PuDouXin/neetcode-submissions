func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1{
		return false
	}
    visited := make(map[int]bool)
	var dfs func(int, int) bool
	dfs = func(i int, p int)bool{
		if visited[i] {
			return false
		}
		visited[i] = true
		for _, edge := range edges{
			if edge[0] == i && edge[1]!=p{
				if !dfs(edge[1],i){
					return false
				}
			}else if edge[1] == i && edge[0]!=p{
				if !dfs(edge[0],i){
					return false
				}
			}
		}
		return true
	}

	return dfs(0,-1) && len(visited) == n
}
