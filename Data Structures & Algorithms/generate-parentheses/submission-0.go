func generateParenthesis(n int) []string {
	
	res := []string{}
	current:=[]string{}
	var backpacking func(int, int)
	backpacking = func(open, closed int){
		if open == n && closed == n{
			res = append(res, strings.Join(current, ""))
			return
		}

		if open < n{
			current = append(current, "(")
			backpacking(open+1,closed)
			current = current[:len(current)-1]
		}
		if closed < open{
			current = append(current, ")")
			backpacking(open,closed+1)
			current = current[:len(current)-1]
		}
	}

	backpacking(0,0)
	return res

}
