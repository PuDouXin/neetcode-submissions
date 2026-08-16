func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
	q := [][2]int{}
	res := 0
	fresh := 0

	dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}

	for i:=0; i< rows; i++{
		for j:=0; j< cols; j++{
			if grid[i][j] == 1{
				fresh++
			}
			if grid[i][j] == 2{
				q=append(q,[2]int{i,j})
			}
		}
	}
	

	for fresh >0 && len(q)>0{
		q2 := [][2]int{}
		for _, node := range q{
			for _, dir := range dirs{
				r,c := node[0]+dir[0], node[1]+dir[1]
				if r >= rows|| c >= cols || r<0 || c<0 || grid[r][c]!=1{
					continue
				}
				grid[r][c]=2
				q2 = append(q2, [2]int{r,c})
				fresh--
			}
		}
		res ++
		q = q2

	}
	if fresh ==0{
		return res
	}
	return -1
}
