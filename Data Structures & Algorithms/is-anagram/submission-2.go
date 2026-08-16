func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
    hitS, hitT := make(map[rune]int),make(map[rune]int)
    for i, c := range s {

		hitS[c]++
        hitT[rune(t[i])]++
	}
    
    for k,v := range hitS{
        if hitT[k]!=v{
            return false
        }
    }

    return true
}
