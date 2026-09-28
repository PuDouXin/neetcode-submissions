func spiralOrder(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
	var res []int
	visited := make(map[[2]int]bool)
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	i, j, k := 0, 0, 0
	for len(res) < n*m {

		index := [2]int{i, j}
		res = append(res, matrix[i][j])
		visited[index] = true

		n_i := i + dir[k][0]
		n_j := j + dir[k][1]

		if n_i < 0 || n_i >= m || n_j < 0 || n_j >= n || visited[[2]int{n_i, n_j}] {
			k = (k + 1) % 4
		}
		i += dir[k][0]
		j += dir[k][1]
	}
	return res

}
