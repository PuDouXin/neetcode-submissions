type dsu struct{
	parent []int
	rank []int
	count int
}

func NewDSU(n int) *dsu{
	parent := make([]int,n)
	rank := make([]int, n)
	for i := range parent{
		parent[i] = i
		rank[i] = 1
	}
	return &dsu{parent: parent, rank: rank, count: n}
}

func (d *dsu) find(n int)int{
	if d.parent[n] != n{
		return d.find(d.parent[n])
	}
	return d.parent[n]
}

func (d *dsu) union(a, b int) bool{
	ra, rb := d.find(a), d.find(b)
	if ra == rb {
		return false
	}
	//ra is the larger one
	if d.rank[ra] < d.rank[rb]{
		ra, rb = rb, ra
	}

	d.parent[rb] = ra
	d.rank[ra] += d.rank[rb]
	d.count --
	return true

}

func minCostConnectPoints(points [][]int) int {
    n := len(points)
	dsu := NewDSU(n)
	var edges [][]int
	for i := 0; i < n; i++{
		x1, y1 := points[i][0], points[i][1]
		for j := i+1; j < n; j++{
			x2, y2 := points[j][0], points[j][1]
			dist := int(math.Abs(float64(x1-x2)))+ int(math.Abs(float64(y1-y2)))
			edges = append(edges,[]int{dist, i, j})
		}

	}
	sort.Slice(edges, func(a, b int) bool{
		return edges[a][0] < edges[b][0]
	})

	res := 0
	for _, edge := range edges{
		dist, a, b := edge[0], edge[1], edge[2]
		if dsu.union(a,b){
			res += dist
		}
	}
	return res
}
