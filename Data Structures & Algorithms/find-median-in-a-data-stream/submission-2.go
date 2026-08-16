type Heap struct {
	data []int
}

func NewHeap() Heap {
	return Heap{data: []int{}}
}

func (m *Heap) Length() int {
	return len(m.data)
}
func (m *Heap) Top() int {

	return m.data[0]
}
func (m *Heap) Push(i int, t int) {
	m.data = append(m.data, i)
	m.SiftUp(len(m.data)-1, t)
}

func (m *Heap) Pop(t int) int {
	top := m.data[0]
	last := len(m.data) - 1
	m.data[0] = m.data[last]
	m.data = m.data[:last]
	m.SiftDown(0, t)
	return top
}

func (m *Heap) SiftUp(i int, t int) { //0 is minHeap, 1 is MaxHeap
	for i > 0 {
		parent := (i - 1) / 2
		if t == 0 {
			if m.data[i] >= m.data[parent] {
				break
			}
		} else {
			if m.data[i] <= m.data[parent] {
				break
			}
		}
		m.data[i], m.data[parent] = m.data[parent], m.data[i]
		i = parent

	}
}

func (m *Heap) SiftDown(i int, t int) { //0 is minHeap, 1 is MaxHeap
	n := len(m.data)
	for {
		left := i*2 + 1
		right := i*2 + 2
		most := i
		if t == 0 { //minHeap
			if left < n && m.data[most] > m.data[left] {
				most = left
			}
			if right < n && m.data[most] > m.data[right] {
				most = right
			}
		} else {
			if left < n && m.data[most] < m.data[left] {
				most = left
			}
			if right < n && m.data[most] < m.data[right] {
				most = right
			}
		}
		if most == i {
			break
		}
		m.data[i], m.data[most] = m.data[most], m.data[i]
		i = most
	}
}

type MedianFinder struct {
	minHeap Heap
	maxHeap Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: NewHeap(),
		maxHeap: NewHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Length() == 0 || num > this.minHeap.Top() {
		this.minHeap.Push(num, 0)
	} else {
		this.maxHeap.Push(num, 1)
	}
	if this.minHeap.Length() > this.maxHeap.Length() && this.minHeap.Length()-this.maxHeap.Length() > 1 {
		t := this.minHeap.Pop(0)
		this.maxHeap.Push(t, 1)
	} else if this.minHeap.Length() < this.maxHeap.Length() && this.maxHeap.Length()-this.minHeap.Length() > 1 {
		t := this.maxHeap.Pop(1)
		this.minHeap.Push(t, 0)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	l := this.minHeap.Length()
	r := this.maxHeap.Length()
	if l == r {
		if l == 0 {
			return 0
		}
		return float64(this.minHeap.Top()+this.maxHeap.Top()) / float64(2)
	}
	if l > r {
		return float64(this.minHeap.Top())
	}
	return float64(this.maxHeap.Top())
}

