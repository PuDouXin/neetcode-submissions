func sortString(s string ) string{
    c := []rune(s)
    sort.Slice(c,func(i, j int) bool{
        return c[i]<c[j]
    })
    return string(c)
}

func groupAnagrams(strs []string) [][]string {

  countS := make(map[string][]string)
  for _,s := range strs{
    sortedS := sortString(s)
    countS[sortedS]= append(countS[sortedS],s)
  }
  var res [][]string
  for _,v:=range countS{
        res = append(res,v)
  }
  return res
}
