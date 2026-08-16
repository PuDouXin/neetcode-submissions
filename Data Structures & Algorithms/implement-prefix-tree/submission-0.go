type TrieNode struct{
	childs [26]*TrieNode
	end	   bool
}
type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: &TrieNode{},
	}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for _, c := range word{
		i := c - 'a'
		if cur.childs[i] == nil{
			cur.childs[i] = &TrieNode{}
		}
		cur = cur.childs[i]
	}
	cur.end = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for _, c := range word{
		i := c - 'a'
		if cur.childs[i] == nil{
			return false
		}
		cur = cur.childs[i]
	}
	return cur.end
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix{
		i := c - 'a'
		if cur.childs[i] == nil{
			return false
		}
		cur = cur.childs[i]
	}
	return true
}
