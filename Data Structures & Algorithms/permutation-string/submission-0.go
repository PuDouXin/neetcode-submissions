func checkInclusion(s1 string, s2 string) bool {
	if len(s1)>len(s2){
		return false
	}

	charS1 := make(map[byte]int)
	for i:=0;i<len(s1);i++{
		charS1[s1[i]]++
	}
	
	
	need := len(charS1)
	for i:=0;i<len(s2);i++{
		charS2 := make(map[byte]int)
		curr :=0
		for j:=i;j<len(s2);j++{
			charS2[s2[j]]++
			if charS2[s2[j]]>charS1[s2[j]]{
				break
			}
			if charS2[s2[j]]==charS1[s2[j]]{
				curr++
			}
		}
		if curr==need{
			return true
		}
	}
	return false
}
