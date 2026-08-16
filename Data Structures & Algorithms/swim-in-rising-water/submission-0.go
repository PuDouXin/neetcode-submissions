type Node struct{
	time, r, c int
}

type MinHeap []Node

func(h MinHeap) Len()int{
	return len(h)
}

func(h MinHeap)Less(i, j int)bool{
	return h[i].time < h[j].time
}

func(h MinHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func(h *MinHeap)Push(x interface{}){
	*h = append(*h, x.(Node))
}

func(h *MinHeap)Pop()interface{}{
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[:n-1]
	return x

}
func swimInWater(grid [][]int) int {
	n := len(grid)
	dirt := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}

	pq := &MinHeap{}
	heap.Init(pq)

	heap.Push(pq, Node{grid[0][0], 0,0})
	visited := make(map[[2]int]bool)
	visited[[2]int{0,0,}]= true

	for pq.Len()>0{
		node := heap.Pop(pq).(Node)
		t,r,c := node.time, node.r, node.c

		if r ==n-1 && c == n-1{
			return t
		}

		for _, dir := range dirt{
			neiR, neiC := r+dir[0], c+dir[1]
			if neiR <0 || neiC <0 || neiR >=n || neiC >=n || visited[[2]int{neiR, neiC}]{
				continue
			}
			visited[[2]int{neiR, neiC}] = true
			heap.Push(pq,Node{max(t, grid[neiR][neiC]),neiR, neiC})
		}
	}
   return -1 
}

func max(a, b int)int{
	if a>=b{
		return a
	}
	return b
}
