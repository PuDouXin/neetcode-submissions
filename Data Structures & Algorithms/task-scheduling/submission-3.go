type task struct {
	f        int
	nextTime int
}

type maxHeap struct {
	data []task
	n    int
	q    []task
}

func NewHeap(n int) maxHeap {
	return maxHeap{data: []task{}, n: n, q: []task{}}
}

func (m *maxHeap) Push(t task) {
	m.data = append(m.data, t)
	m.SiftUp(len(m.data) - 1)
}

func (m *maxHeap) Pop() task {
	top := m.data[0]
	last := len(m.data) - 1
	m.data[0] = m.data[last]
	m.data = m.data[:last]
	m.SiftDown(0)

	return top
}

func (m *maxHeap) SiftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if m.data[parent].f >= m.data[i].f {
			break
		}
		m.data[parent], m.data[i] = m.data[i], m.data[parent]
		i = parent
	}
}

func (m *maxHeap) SiftDown(i int) {
	length := len(m.data)
	for {
		l := i*2 + 1
		r := i*2 + 2
		largest := i
		if l < length && m.data[l].f > m.data[largest].f {
			m.data[l], m.data[largest] = m.data[largest], m.data[l]
		}
		if r < length && m.data[r].f > m.data[largest].f {
			m.data[r], m.data[largest] = m.data[largest], m.data[r]
		}
		if largest == i {
			break
		}
		i = largest
	}
}

func leastInterval(tasks []byte, n int) int {
	frequency := make(map[byte]int)
	for _, b := range tasks {
		frequency[b]++
	}
	m := NewHeap(n)
	for _, v := range frequency {
		m.Push(task{f: v, nextTime: 0})
	}

	time := 0
	for len(m.data) > 0 || len(m.q) > 0 {
		time++
		if len(m.data) == 0 {
			time = m.q[0].nextTime
		} else {
			top := m.Pop()
			newF := top.f - 1
			if newF > 0 {
				m.q = append(m.q, task{f: newF, nextTime: n + time})
			}
		}
		if len(m.q) > 0 && m.q[0].nextTime == time {
			m.Push(m.q[0])
			m.q = m.q[1:]
		}

	}
	return time
}