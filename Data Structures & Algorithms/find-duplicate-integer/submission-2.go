func abs(num int)int{
	if num < 0 {
		return -num
	}
	return num
}

func findDuplicate(nums []int) int {
    for _,n := range nums{
		idx := abs(n)-1
		if nums[idx]<0{
			return abs(n)
		}
		nums[idx] *= -1
	}
	return -1
}
