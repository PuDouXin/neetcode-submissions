func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(a,b int)bool{
		return intervals[a][0]<intervals[b][0]
	})
	n := len(intervals)
	i := 0
	res := 0
	preEnd := math.MinInt
	for i<n{
		if intervals[i][0]>=preEnd{
			preEnd = intervals[i][1]
		}else{
			res ++
			preEnd = min(intervals[i][1],preEnd)
		}
		i++
	}
	return res

}
