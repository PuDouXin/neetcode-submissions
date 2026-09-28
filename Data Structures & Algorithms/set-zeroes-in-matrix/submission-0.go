func setZeroes(matrix [][]int) {
    m, n := len(matrix), len(matrix[0])
	r := make([]int,m)
	c := make([]int,n)

	for i := 0; i<m; i++{
		for j:= 0;j<n;j++{
			if matrix[i][j] ==0{
				r[i] = 1
				c[j] =1
			}
		}
	}

	for i := 0; i<m; i++{
		for j:= 0;j<n;j++{
			if r[i]==1 || c[j] == 1{
				matrix[i][j] =0
			}
		}
	}
}
