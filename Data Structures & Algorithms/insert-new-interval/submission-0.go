func insert(intervals [][]int, newInterval []int) [][]int {
	i :=0
	n := len(intervals)
	var res [][]int
    for i<n &&intervals[i][1]<newInterval[0]{
		
		res = append(res,intervals[i])
		
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1]{
		newInterval[0] = min(intervals[i][0],newInterval[0])
		newInterval[1] = max(intervals[i][1],newInterval[1])
		i++
	}
	res=append(res,newInterval)
	for i<n{
		res = append(res,intervals[i])
		i++
	}
	return res
	
}
