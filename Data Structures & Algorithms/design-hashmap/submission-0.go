type Node struct{
	key,val int
	next *Node
}
type MyHashMap struct {
	data []*Node
}

func Constructor() MyHashMap {
    data := make([]*Node,1000)
	for i := range data{
		data[i] = &Node{key:-1,val:-1}
	}
	return MyHashMap{data: data}
}

func (this *MyHashMap) hash(key int) int{
	return key%len(this.data)
}

func (this *MyHashMap) Put(key int, value int) {
    cur:=this.data[this.hash(key)]
	for cur.next!=nil{
		if cur.next.key==key{
			cur.next.val=value
			return
		}
		cur = cur.next
	}
	cur.next = &Node{key: key,val: value}
}

func (this *MyHashMap) Get(key int) int {
    cur:=this.data[this.hash(key)]
	for cur.next!=nil{
		if cur.next.key==key{
			return cur.next.val
		}
		cur = cur.next
	}
	return -1

}

func (this *MyHashMap) Remove(key int) {
	cur:=this.data[this.hash(key)]
    for cur.next!=nil{
		if cur.next.key==key{
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */