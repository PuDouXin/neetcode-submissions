func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(a,b int)bool{
		return intervals[a][0]<intervals[b][0]
	})
	n := len(intervals)
	var res [][]int
	res = append(res, intervals[0])
	i :=0
	j := 0
	for i<n{
		
		if intervals[i][0] > res[j][1]{
			res = append(res,intervals[i])
			j++
		}else{
			s := min(res[j][0],intervals[i][0])
			e := max(res[j][1],intervals[i][1])
			res[j][0] = s
			res[j][1] =e
		}
		i++
	}
	return res
}
