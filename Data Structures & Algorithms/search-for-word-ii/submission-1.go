type TrieNode struct {
	childs [26]*TrieNode
	idx     int
	ref		int
}
type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{idx:-1}}
}

func (this *WordDictionary) AddWord(word string, j int) {
	cur := this.root
	for _, c := range word {
		i := c - 'a'
		if cur.childs[i] == nil {
			cur.childs[i] = &TrieNode{idx: -1}
		}
		cur = cur.childs[i]
		cur.ref++
	}
	cur.idx = j

}

func findWords(board [][]byte, words []string) []string {
	tree := Constructor()
	for i, w := range words {
		tree.AddWord(w, i)
	}
	res := []string{}
	lenR, lenC := len(board), len(board[0])

	var dfs func(r int, c int, node *TrieNode)
	dfs = func(r int, c int, node *TrieNode) {

		if r < 0 || c < 0 || r >= lenR || c >= lenC || board[r][c] == '*' || node.childs[board[r][c]-'a'] == nil {
			return
		}
		ch := board[r][c]
	
		child := node.childs[ch-'a']
		board[r][c] = '*'
		if child.idx != -1 {
			res = append(res, words[child.idx])
			child.idx = -1
			child.ref--
			if child.ref == 0 {
				node.childs[ch-'a'] = nil
				board[r][c] = ch
				return
			}
		}
		dfs(r+1, c, child)
		dfs(r, c+1, child)
		dfs(r-1, c, child)
		dfs(r, c-1, child)
		board[r][c] = ch

	}
	for r := 0; r < lenR; r++ {
		for c := 0; c < lenC; c++ {
			dfs(r, c, tree.root)
		}
	}

	return res
}