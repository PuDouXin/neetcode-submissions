type Node struct {
	key int
	value int
	pre,next *Node
}
type LRUCache struct {
	capacity int
	cache map[int]*Node
    left,right *Node
}

func Constructor(capacity int) LRUCache {
    lru :=LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		left: &Node{},
		right: &Node{},
	}
	lru.left.next = lru.right
	lru.right.pre = lru.left
	return lru
}

func (this *LRUCache) remove( node *Node){
	p := node.pre
	n := node.next
	p.next = n
	n.pre = p
}

func (this *LRUCache)insert(node *Node){
	p := this.right.pre
	p.next = node
	node.next=this.right
	node.pre = p
	this.right.pre = node
}

func (this *LRUCache) Get(key int) int {
    if v,exist := this.cache[key]; exist{
		this.remove(v)
		this.insert(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if v,exist := this.cache[key]; exist{
		this.remove(v)
		delete(this.cache,key)
	}

	n:=&Node{
			key: key,
			value: value,
		}
		this.cache[key]=n
		this.insert(n)
		if len(this.cache)>this.capacity{
			ln := this.left.next
			this.remove(ln)
			delete(this.cache,ln.key)
		}
}
