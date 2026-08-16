func findWords(board [][]byte, words []string) []string {
    res := []string{}
	lenL, lenR := len(board), len(board[0])
	
	var backpacking func(int, int,int, string)bool
	backpacking = func(r,c,i int, word string)bool{
		if i==len(word){
			return true
		}
		if r<0 || c<0 || r>=lenL || c >= lenR || board[r][c] != word[i]{
			return false
		}
		board[r][c]=1
		result := backpacking(r+1, c, i+1, word) ||
			  backpacking(r, c+1, i+1, word) ||
			  backpacking(r-1, c, i+1, word) ||
			  backpacking(r, c-1, i+1, word)
		board[r][c]= word[i]
		return result
	}
	for _, w := range words{
		found := false
		for r:= 0; r<lenL; r++{
			if found{
				break
			}
			
			for c :=0;c<lenR; c++{
				if board[r][c] == w[0] && backpacking(r,c,0,w){
					res = append(res, w)
					found = true
					break
				}
			}
		}
	}
	return res
}
