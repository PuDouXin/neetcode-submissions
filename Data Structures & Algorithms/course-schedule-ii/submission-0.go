func findOrder(numCourses int, prerequisites [][]int) []int {
    indegree := make([]int, numCourses)
	adj := make([][]int, numCourses)

	for _, preq := range prerequisites{
		src, dst := preq[0], preq[1]
		indegree[dst]++
		adj[src]= append(adj[src],dst)
	}
	q := []int{}
	for i:=0; i <numCourses; i++{
		if indegree[i] ==0{
			q = append(q, i)
		}
	}
	output := []int{}
	finish := 0

	for len(q) >0{
		node := q[0]
		q = q[1:]
		output = append([]int{node}, output...)
		finish ++
		for _, a := range adj[node]{
			indegree[a]--
			if indegree[a] ==0{
				q = append(q, a)
			}
		}
	}
	if finish != numCourses{
		return []int{}
	}
	return output
}
