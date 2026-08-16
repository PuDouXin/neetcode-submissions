func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
    res := make([][]int,0)

	dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}

	pacific := make(map[[2]int]bool)
	atlantic := make(map[[2]int]bool)
	var dfs func(i,j int, visit map[[2]int]bool, prevHeight int)
	dfs = func(i int, j int, visit map[[2]int]bool,prevHeight int){
		coord := [2]int{i,j}
		if i>=rows || j>= cols || i<0 || j<0 || visit[coord] || heights[i][j] < prevHeight{
			return
		}
		visit[coord] = true
		for _, dir := range dirs{
			dfs(i+dir[0],j+dir[1],visit,heights[i][j])
		}

	}

	for c := 0;c < cols;c++{
		dfs(0,c, pacific, heights[0][c])
		dfs(rows-1,c, atlantic, heights[rows-1][c])
	}

	for r := 0; r< rows; r++{
		dfs(r,0,pacific, heights[r][0])
		dfs(r, cols-1, atlantic, heights[r][cols-1])
	}

	for key,_ := range pacific{
		if atlantic[key]{
			res = append(res,[]int{key[0],key[1]})
		}
	}
	return res
}
