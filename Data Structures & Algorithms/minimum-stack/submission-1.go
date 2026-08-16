type MinStack struct {
  value []int
  min int
}

func Constructor() MinStack {
	return MinStack{
		value: []int{},
		min: math.MaxInt64,
	}

}

func (this *MinStack) Push(val int) {
	if len(this.value)==0{
		this.value = append(this.value,0)
		this.min = val
	}else{
		this.value = append(this.value,val-this.min)
		if val<this.min{
			this.min=val
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.value)==0{
		return
	}
	res := this.value[len(this.value)-1]
	this.value=this.value[:len(this.value)-1]
	if res<0{
		this.min=this.min-res
	}
}

func (this *MinStack) Top() int {
	val := this.value[len(this.value)-1]
	if val>0{
		return this.value[len(this.value)-1]+this.min
	}
	return this.min
	
}

func (this *MinStack) GetMin() int {
	return this.min
}
