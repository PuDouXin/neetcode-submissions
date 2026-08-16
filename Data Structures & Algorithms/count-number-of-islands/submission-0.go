func numIslands(grid [][]byte) int {
    lenR, lenC := len(grid), len(grid[0])
	res := 0
	var dfs func(int, int, [][]byte)
	dfs = func (i, j int, grid [][]byte){
		if i<0 || j<0 || i >= lenR || j >= lenC || grid[i][j] == '0'|| grid[i][j] == '*'{
			return
		}
		

		grid[i][j] = '*'

		dfs(i+1, j, grid)
		dfs(i, j+1, grid)
		dfs(i-1, j, grid)
		dfs(i, j-1, grid)
		
	}

	for i := 0; i< lenR; i++{
		for j := 0; j < lenC; j++ {
			if grid[i][j] == '1'{
				dfs(i,j, grid)
				res++
			}
		}
	}

	return res
}
