type edge struct{
	v int
	t int
}
func networkDelayTime(times [][]int, n int, k int) int {
    adj := make(map[int][]edge)

	for _, t := range times{
		u,v,t := t[0], t[1], t[2]
		adj[u] = append(adj[u],edge{v:v, t:t})
	}

	distance := make([]int,n)
	for i := range distance{
		distance[i] = math.MaxInt32
	}

	var dfs func(int,int)
	dfs = func(node int, time int){
		if time >=distance[node-1]{
			return
		}
		distance[node-1] = time
		for _, e := range adj[node]{
			next, w := e.v, e.t
			dfs(next, w+time)
		}
	}

	dfs(k,0)
	res := 0
	for _, dist := range distance{
		if dist == math.MaxInt32{
			return -1
		}
		if dist > res{
			res = dist
		}
	}
	return res
}
