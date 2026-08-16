type minHeap struct{
	data []int
	k int
}

func NewHeap(k int) minHeap{
	return minHeap{ data: []int{}, k: k}
}

func (m *minHeap)Push(i int){
	m.data = append(m.data, i)
	m.SiftUp(len(m.data)-1)
	if len(m.data)>m.k{
		m.Pop()
	}
}

func (m *minHeap)Pop() int{
	if len(m.data) ==0{
		return -1
	}
	top := m.data[0]
	last := len(m.data) -1
	m.data[0] = m.data[last]
	m.data = m.data[:last]
	m.SiftDown(0)
	return top

}

func (m *minHeap)SiftDown(i int){
	n := len(m.data) 
	for{
		left := i*2 + 1
		right := i*2 + 2
		smallest := i
		if left < n && m.data[left] < m.data[smallest]{
			smallest = left
		}

		if right < n && m.data[right] < m.data[smallest]{
			smallest = right
		}
		if smallest == i {
			break
		}
		m.data[smallest], m.data[i] = m.data[i], m.data[smallest]
		i = smallest
	}
}

func (m *minHeap)SiftUp(i int){
	for i >0{
		parent := (i-1)/2
		if m.data[parent] <= m.data[i]{
			break
		}
		m.data[parent], m.data[i] =  m.data[i], m.data[parent]
		i = parent
	}
}

func findKthLargest(nums []int, k int) int {
	m := NewHeap(k)
	for _, n := range nums{
		m.Push(n)
	}
	return m.data[0]
}
