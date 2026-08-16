func characterReplacement(s string, k int) int {
	frequency := make(map[byte]int)
	for i:=0;i<len(s);i++{
		frequency[s[i]]++
	}
	res := 0
	for c := range frequency{
		l, count := 0,0
		for i:=0;i<len(s);i++{
			if s[i]==c{
				count++
			}

			for i-l+1 - count >k{
				if s[l] == c{
					count --
				}
				l++
			}

			if i-l+1> res{
				res = i-l+1
			}
		}
	}

	return res

}
