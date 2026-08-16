func isPalindrome (s string,l, r int)bool{
	
	for l < r{
		if s[l]!=s[r]{
			return false
		}
		l++
		r--
	}
	return true
}
func partition(s string) [][]string {
	res := [][]string{}
	current := []string{}
	var backpacking func(int)
	backpacking = func (i int){
		if i >=len(s){
			res = append(res, append([]string{},current...))
		}
		for j := i; j<len(s);j++{
			if isPalindrome(s, i, j){
				current = append(current, s[i:j+1])
				backpacking(j+1)
				current = current[:len(current)-1]
			}
		}

	}
	backpacking(0)
	return res
	
}
