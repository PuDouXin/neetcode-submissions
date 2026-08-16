func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	chars := map[byte]string{
		'2' : "abc",
		'3' : "def",
		'4' : "ghi",
		'5' : "jkl",
		'6' : "mno",
		'7' : "pqrs",
		'8' : "tuv",
		'9' : "wxyz",
	}
	res := []string{}
	

	var backpacking func(int, string)
	backpacking = func(i int, curStr string){
		if i == len(digits){
			res = append(res, curStr)
			return
		}
		for _ , c := range chars[digits[i]]{
			backpacking(i+1, curStr + string(c))
		}
	}
	backpacking(0, "")
	return res
}
