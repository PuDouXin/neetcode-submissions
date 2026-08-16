func longestCommonPrefix(strs []string) string {
    s := strs[0]
    for i:=1;i<len(strs);i++{
        j:=0
        for j<len(s)&& j<len(strs[i]){
            if s[j]!=strs[i][j]{
                break
            }
            j++
        }
        s = s[:j]
    }
    return s
}
