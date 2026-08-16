func containsT(a map[byte]int, b map[byte]int)bool{
	for k,v := range a{
		if b[k]<v{
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	if len(s)<len(t){
		return ""
	}
    charSet := make(map[byte]int)
	for i:=0;i<len(t);i++{
		charSet[t[i]]++
	}
	count := 1000
	var res string
	
	for l:=0;l<len(s);l++{
		char := make(map[byte]int)
		for r:=l;r<len(s);r++{
			char[s[r]]++
			if containsT(charSet,char){
				if r-l+1<count{
					count = r-l+1
					res = s[l:r+1]
				}
			}
		}
	}
	return res
}
