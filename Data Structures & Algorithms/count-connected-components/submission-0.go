func countComponents(n int, edges [][]int) int {
    
	visited := make(map[int]bool)
	res := 0

	var dfs func(int, int)
	dfs = func(i int, p int){
		if visited[i]{
			return
		}

		visited[i] = true

		for _, edge := range edges{
			if edge[0] == i && edge[1] != p{
				dfs(edge[1], i)
			}else if edge[1] == i && edge[0] != p{
				dfs(edge[0], i)
			}
		}
	}
	for i:=0; i< n ;i++{
		if !visited[i]{
			dfs(i, -1)
			res++
		}
	}
	return res
}
