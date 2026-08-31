func longestIncreasingPath(matrix [][]int) int {
	m,n := len(matrix),len(matrix[0])
    memo := make([][]int,m)
	for i := range memo{
		memo[i] = make([]int, n)
	}
	dirs := [][2]int{{0,1},{1,0},{0,-1},{-1,0}}

	var dp func(int, int, int)int

	dp = func(r,c,pre int)int{
		if r>=m || r<0 || c>=n || c<0|| matrix[r][c]<=pre{
			return 0
		}
		if memo[r][c]!=0{
			return memo[r][c]
		}
		res := 1
		for i := range dirs{
			res = max(res,1+dp(r+dirs[i][0],c+dirs[i][1],matrix[r][c]))
		}
		memo[r][c] = res
		return res
	}
	res := 1
	for r := range m{
		for c := range n{
			res = max(res, dp(r,c,-1))
		}
	}
	return res
}
