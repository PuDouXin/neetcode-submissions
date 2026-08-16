type maxHeap struct{
	data []int 
}

func NewHeap() maxHeap{
	return maxHeap{data: []int{}}
}

func (m *maxHeap)Push(i int){
	m.data = append(m.data, i)
	m.SiftUp(len(m.data)-1)
}

func (m *maxHeap)Pop()int{
	if len(m.data) == 0{
		return -1
	}
	top := m.data[0]
	last := len(m.data)-1
	m.data[0] = m.data[last]
	m.data = m.data[:last]
	if len(m.data) > 0{
		m.SiftDown(0)
	}
	return top
}

func (m *maxHeap)SiftUp(i int){
	
	for i > 0{
		parent := (i-1)/2
		if m.data[parent]>=m.data[i]{
			break
		}
		m.data[parent], m.data[i] = m.data[i], m.data[parent]
		i = parent
	}
}

func (m *maxHeap)SiftDown(i int){
	n := len(m.data)
	for {
		left := i*2+1
		right := i*2+2
		largest := i
		if left < n && m.data[largest] < m.data[left]{
			largest = left
		}
		if right < n && m.data[largest] < m.data[right]{
			largest = right
		}
		if largest == i{
			break
		}
		m.data[i], m.data[largest] = m.data[largest],m.data[i]
		i = largest

	}
}

func lastStoneWeight(stones []int) int {
	h := NewHeap()
	for _, v := range stones{
		h.Push(v)
	}
	x , y := 0 , 0
	for len(h.data)>1{
		x = h.Pop()
		y = h.Pop()
		if y < x{
			h.Push(x-y)
		}
	}
	if len(h.data) == 0{
		return 0
	}
	return h.data[0]

}
