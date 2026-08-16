

func largestRectangleArea(heights []int) int {
	var stack []int
	res:=0
	for i:=0;i<=len(heights);i++{
		currentH:=0
		if i<len(heights){
			currentH = heights[i]
		}
		for len(stack)>0 && heights[stack[len(stack)-1]]>currentH{
			topIndex :=stack[len(stack)-1]
			
			stack = stack[:len(stack)-1]
			height := heights[topIndex]
			w := i
			if len(stack)>0{
				left := stack[len(stack)-1]
				w = i-left-1
			}
			aero := w*height
			if aero>res{
				res= aero
			}
		}
		stack = append(stack,i)
	}
	return res
}
