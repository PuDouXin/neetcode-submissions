func findDuplicate(nums []int) int {
    dup := make(map[int]bool)
	for _,n:= range nums{
		if dup[n]{
			return n
		}
		dup[n]=true
	}
	
	return -1
}
