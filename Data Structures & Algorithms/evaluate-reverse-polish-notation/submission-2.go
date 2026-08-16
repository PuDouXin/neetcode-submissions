type stack []string

func(s *stack) Pop() string{
	index := len(*s)-1
	v:=(*s)[index]
	*s = (*s)[:index]
	return v
} 

func(s *stack) Push(v string){
	*s = append(*s,v)
}

func evalRPN(tokens []string) int {
	var st stack
	for _,s:=range tokens{
		var res int
		if s == "*" || s=="+" || s=="/" || s=="-"{
			a,_ := strconv.Atoi(st.Pop())
			b,_ := strconv.Atoi(st.Pop())
			switch s{
				case "*":
					res =a*b
				case "/":
					res = b/a
				case "+":
					res = a+b
				case "-":
					res = b-a
			}
			st.Push(strconv.Itoa(res))
		}else{
			st.Push(s)
		}
		
	}
	r,_:=strconv.Atoi(st.Pop())
	return r
}
