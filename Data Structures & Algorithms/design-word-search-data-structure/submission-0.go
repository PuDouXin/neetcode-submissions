type TrieNode struct{
	childs [26]*TrieNode
	end bool
}
type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{root: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string)  {
    cur := this.root
	for _,c := range word{
		i := c - 'a'
		if cur.childs[i] == nil{
			cur.childs[i] = &TrieNode{}
		}
		cur = cur.childs[i]
	}
	cur.end = true
}

func (this *WordDictionary) dfs(cur *TrieNode, word string, i int) bool {
	if i == len(word) {
		return cur.end
	}
	c := word[i]
	if c == '.' {
		for _, next := range cur.childs {
			if next != nil && this.dfs(next, word, i+1) {
				return true
			}
		}
	} else {
		index := c - 'a'
		if cur.childs[index] != nil && this.dfs(cur.childs[index], word, i+1) {
			return true
		}
	}
	return false
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(this.root, word, 0)
}