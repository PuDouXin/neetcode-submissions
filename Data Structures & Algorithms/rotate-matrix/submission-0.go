func rotate(matrix [][]int)  {

    n := len(matrix)
	rotation := make([][]int, n)
	for i := range matrix{
		rotation[i] = make([]int, n)
	}
	for i := range matrix{
		for j := range matrix[0]{
			rotation[j][n-1-i] =  matrix[i][j]
		}
	}

	for i := range matrix{
		for j := range matrix[0]{
			 matrix[i][j] = rotation[i][j]
		}
	}
}
