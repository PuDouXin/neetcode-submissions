func solve(board [][]byte) {
    rows, cols := len(board), len(board[0])
	dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}

	var dfs func(int, int)
	dfs = func(i, j int){
		if i >= rows || j >= cols || i<0 ||j<0 || board[i][j] != 'O'{
			return
		}
		board[i][j] = '#'
		for _, dir := range dirs{
			dfs(i+dir[0], j+dir[1])
		}
	}

	for i := 0; i< rows; i++{
		if board[i][0] == 'O'{
			dfs(i,0)
		}
		if board[i][cols-1] == 'O'{
			dfs(i, cols -1)
		}
	}

	for i := 0; i< cols; i++{
		if board[0][i] == 'O'{
			dfs(0, i)
		}
		if board[rows - 1][i] == 'O'{
			dfs(rows - 1, i)
		}
	}

	for  i := 0; i< rows; i++{
		for j :=0; j<cols; j++{
			if board[i][j] ==  'O'{
				board[i][j] = 'X'
			}
			
		}
	}

	for  i := 0; i< rows; i++{
		for j :=0; j<cols; j++{
			
			if board[i][j] ==  '#'{
				board[i][j] = 'O'
			}
		}
	}
}
