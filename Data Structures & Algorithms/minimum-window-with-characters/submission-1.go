

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
	need:= len(charSet)
	
	l:=0
		char := make(map[byte]int)
		have:=0
		for r:=0;r<len(s);r++{
			char[s[r]]++
			if charSet[s[r]]>0 && char[s[r]]==charSet[s[r]]{
				have++
			}
			for have==need{
				if r-l+1<count{
					count = r-l+1
					res = s[l:r+1]
				}
				char[s[l]]--
				if char[s[l]]<charSet[s[l]]{
					have--
				}
				l++
			}
			
		}
	
	return res
}
