func partitionLabels(s string) []int {
    lastIdx := make(map[rune]int)
	for i,c:= range s{
		lastIdx[c] = i
	}
	var res []int
	end,size := 0,0
	for i,c := range s{
		size++
		if lastIdx[c]>end{
			end = lastIdx[c]
		}
		if i ==end{
			res = append(res,size)
			size =0
		}
	}
	return res
}
