type stack []rune

func(s *stack)Push(a rune){
	*s=append(*s, a)
}

func(s *stack)Pop() rune{
	index := len(*s)-1
	if index<0{
		return 'a'
	}
  	res := (*s)[index]
  	*s = (*s)[:index]
 	return res
}

func isValid(s string) bool {
	pair := map[rune]rune{']':'[','}':'{',')':'('}
	var st stack
    for _, char := range s{
		switch rune(char){
			case '[','(','{': 
				st.Push(char)
			case ']','}',')':
				v:= st.Pop()
				if v!=pair[char]{
					return false
				}
				
		}
	}
	if len(st)>0{
		return false
	}
	return true
}
