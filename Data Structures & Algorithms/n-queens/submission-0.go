func isSafe(r int, c int, board [][]string)bool{
	for row := r-1; row >=0; row --{
		if board[row][c] == "Q"{
			return false
		}
	}
	for row, col := r-1, c-1; row >=0 && col >= 0; row, col = row-1, col-1{
		if board[row][col] == "Q"{
			return false
		}
	}
	for row, col := r-1, c+1; row >=0 && col < len(board); row, col = row-1, col+1{
		if board[row][col] == "Q"{
			return false
		}
	}
	return true
}
func solveNQueens(n int) [][]string {
	res := [][]string{}
	board := make([][]string, n)
	for i := range board{
		board[i] = make([]string, n)
		for j := range board[i]{
			board[i][j] = "."
		}
	}
	var backpacking func(int)

	backpacking = func(i int){
		if i == n{
			tmp := make([]string,n)
			for i := range board{
				tmp[i] = ""
				for j := range board[i]{
					tmp[i] += board[i][j]
				}
			}
			res = append(res,tmp)
		}
		for j := 0;j< n ;j++{
			if isSafe(i,j, board){
				board[i][j] = "Q"
				backpacking(i+1)
				board[i][j] = "."
			}
		}
	}
	backpacking(0)
	return res
}
