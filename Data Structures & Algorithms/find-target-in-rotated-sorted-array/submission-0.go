func search(nums []int, target int) int {
res := -1
 for i,v := range nums{
    if v==target{
        return i
    }
 }
 return res
}
