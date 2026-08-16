
func min (a int, b int) int{
	if a>=b{
		return b
	}
	return a
}
func maxArea(heights []int) int {
 res := 0
 start := 0
 end := len(heights) -1
 for start < end {
	m := min(heights[start],heights[end])
	if m*(end-start)>res{
		res = m*(end-start)
	}
	if m == heights[start]{
		start++
	}else{
		end --
	}
 }
 return res
}
