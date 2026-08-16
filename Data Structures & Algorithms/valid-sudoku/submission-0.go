func isValidSudoku(board [][]byte) bool {
 rows := make([]map[byte]bool,9)
 cols := make([]map[byte]bool,9)
 squares := make([]map[byte]bool,9)

 for i := 0; i<9 ;i++{
	rows[i] = make(map[byte]bool)
	cols[i] =  make(map[byte]bool)
	squares[i] =  make(map[byte]bool)
 }
 for i,r := range board{
	for j,c := range r{
		if c =='.' {
			continue
		}
		squareIdx := (i/3)*3+j/3
		if rows[i][c] || cols[j][c] || squares[squareIdx][c]{
			return false
		}
		rows[i][c] = true
		cols[j][c] = true
		squares[squareIdx][c] =true
	}
 }
 return true

}
