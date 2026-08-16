
type item struct {
	v int
	i int
}
type stack []item

func (s *stack)Pop() item{
  res := (*s)[len(*s)-1]
  *s =(*s)[:len(*s)-1]
  return res
}

func (s *stack)Push(v int,i int){
 (*s) = append((*s),item{v:v,i:i})
}

func(s *stack)Top()(bool,item){
	if len(*s)==0{
		return false,item{}
	}
	return true,(*s)[len(*s)-1]
}



func dailyTemperatures(temperatures []int) []int {
	var s stack
	res :=make([]int,len(temperatures))
	for i,t := range temperatures{
			
		for {
			ok, tv := s.Top()
			if !ok || tv.v >= t {
				break
			}
			r := s.Pop()
			res[r.i] = (i - r.i)
		}

			s.Push(t,i)
		
	}
	for len(s)>0{
		r:=s.Pop()
		res[r.i]=0
	}

	return res
}
