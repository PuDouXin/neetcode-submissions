
func min (a int, b int) int{
	if a>=b{
		return b
	}
	return a
}
func maxArea(heights []int) int {
 res := 0
 for i := 0; i< len(heights)-1; i++{
	for j:= i+1; j < len(heights); j++{
		tmp := min(heights[i],heights[j])*(j-i)
		if tmp > res{
			res = tmp
		}
	}
 }
 return res
}
