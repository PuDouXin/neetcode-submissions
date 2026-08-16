func foreignDictionary(words []string) string {
    adj := make(map[byte]map[byte]struct{})
	indegrees := make(map[byte]int)

	//build graph
	for _, word := range words{
		for i := 0; i< len(word); i++{
			char := word[i]
			if _, exists := adj[char]; !exists{
				adj[char] = make(map[byte]struct{})
			}
			indegrees[char] = 0
		}
	}
	for i:=0;i<len(words) -1; i++{
		w1, w2 := words[i], words[i+1]
		minLen := len(w1)
		if len(w2)< minLen{
			minLen = len(w2)
		}
		 
		if len(w1) > len(w2) && w1[:minLen] ==w2[:minLen]{
			return ""
		}
		for j := 0; j<minLen;j++{
			if w1[j]!=w2[j]{
				if _, exists := adj[w1[j]][w2[j]]; !exists{
					adj[w1[j]][w2[j]] = struct{}{}
					indegrees[w2[j]]++
				}
				break
			}
		}
	}

	//topological sort
		q := []byte{}
		for char := range indegrees{
			if indegrees[char] ==0{
				q = append(q,char)
			}
		}

		res := []byte{}
		for len(q)>0{
			char := q[0]
			q = q[1:]
			res = append(res, char)

			for neighbor := range adj[char]{
				indegrees[neighbor] --
				if indegrees[neighbor] ==0{
					q = append(q, neighbor)
				}
			}
		}
		
	
	

	if len(res) != len(indegrees){
			return ""
		}
	return string(res)	
}
