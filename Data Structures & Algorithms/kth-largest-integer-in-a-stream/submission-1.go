type MinHeap struct {
	data []int
}

func NewMinHeap() MinHeap{
	return MinHeap{data: []int{}}
}
func (this *MinHeap)Len()int{
	return len(this.data)
}
func (this *MinHeap)Top() int{
	return this.data[0]
}
func(this *MinHeap)ShiftUp(i int){
	for i>0{
		parent := (i-1)/2
		if this.data[parent]<=this.data[i]{
			break
		}
		this.data[parent],this.data[i] = this.data[i],this.data[parent]
		i = parent
	}
}

func(this *MinHeap)ShiftDown(i int){
	n := len(this.data)
	for{
		left := i*2+1
		right := i*2+2
		smallest := i
		if left < n && this.data[smallest]>this.data[left]{
			smallest = left
		}
		if right < n &&this.data[smallest]>this.data[right]{
			smallest = right
		}
		if i==smallest{
			break
		}
		this.data[smallest],this.data[i] = this.data[i],this.data[smallest]
		i = smallest
	}
}

func (this *MinHeap)Push(i int){
	this.data = append(this.data,i)
	this.ShiftUp(len(this.data)-1)
}
func (this *MinHeap) Pop() int{
	top := this.data[0]
	last := len(this.data)-1
	this.data[0] = this.data[last]
	this.data = this.data[:last]

	this.ShiftDown(0)
	return top
}

type KthLargest struct {
    k int
	minHeap MinHeap
}




func Constructor(k int, nums []int) KthLargest {
	kl :=KthLargest{minHeap: NewMinHeap(), k: k}
	for _, v := range nums{
		kl.Add(v)
	}
	return kl
}


func (this *KthLargest) Add(val int) int {
	if this.minHeap.Len()< this.k{
    	this.minHeap.Push(val)
	}else if val > this.minHeap.Top(){
		this.minHeap.Pop()
		this.minHeap.Push(val)
	}
	return this.minHeap.Top()
}
