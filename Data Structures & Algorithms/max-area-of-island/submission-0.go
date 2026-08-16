func maxAreaOfIsland(grid [][]int) int {
    lenR, lenC := len(grid), len(grid[0])
	res := 0
	var dfs func(int, int, [][]int) int
	dfs = func (i int , j int,  grid [][]int) int{
		if i<0 || j<0 || i >= lenR || j >= lenC || grid[i][j] == 0|| grid[i][j] == -1{
			return 0
		}
		

		grid[i][j] = -1
		

		return 1+dfs(i+1, j, grid)+dfs(i, j+1, grid)+dfs(i-1, j,  grid)+dfs(i, j-1, grid)
		
		
	}

	for i := 0; i< lenR; i++{
		for j := 0; j < lenC; j++ {
			if grid[i][j] == 1{
				s :=dfs(i,j, grid)
				if s > res{
					res = s
				}
			}
		}
	}
	return res
}
