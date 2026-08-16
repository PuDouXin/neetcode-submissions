type pair struct{
	row int
	col int
}
func exist(board [][]byte, word string) bool {
	seen := make(map[pair]bool)
	rows, cols := len(board), len(board[0])
	

	var backpacking func(int, int, int) bool
	backpacking = func(row, col, index int)bool{
		if index == len(word){
			return true
		}
		if row <0 || col < 0||row >= rows || col >= cols || board[row][col] != word[index] || seen[pair{row:row, col:col}]{
			return false
		}
		seen[pair{row:row, col:col}] = true 
		res := backpacking(row+1, col, index+1) ||
				 backpacking(row, col+1, index+1) ||
				  backpacking(row-1, col, index+1) ||
				   backpacking(row, col-1, index+1)
		delete(seen,pair{row:row, col:col})
		return res
	}

	for r:=0; r< rows; r++{
		for c:=0; c<cols; c++{
			if backpacking(r,c,0){
				return true
			}
		}
	}
	return false

}
