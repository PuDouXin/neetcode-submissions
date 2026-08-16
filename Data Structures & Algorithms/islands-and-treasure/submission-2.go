func islandsAndTreasure(grid [][]int) {
  rows, cols := len(grid), len(grid[0])
  q := [][2]int{}
  for i:=0; i<rows; i++{
	for j:= 0; j < cols; j++{
		if grid[i][j] == 0{
			q = append(q, [2]int{i,j})
		}
	}
  }
  if len(q) ==0{
	return 
  }
  dirs := [][2]int{{0,1},{0, -1},{1,0},{-1,0}}
  for len(q) >0{
	node := q[0]
	q = q[1:]
	row,col := node[0], node[1]

	for _, dir := range dirs{
		r, c := row+dir[0], col+dir[1]
		if r >= rows || c >= cols || r<0 || c<0 || grid[r][c] != 2147483647{
			continue
		}
		grid[r][c] = grid[row][col]+1
		q = append(q, [2]int{r,c})
	}
  }
}
