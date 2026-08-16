func binary(nums []int, l int, r int, target int) int{
   
    for l<=r{
        m:= (l+r)/2
        if nums[m]==target{
            return m
        }
        if nums[m]>target{
            r=m-1
        }else{
            l=m+1
        }
    }
    return -1
}



func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		//left part is sorted
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	pivot := l
	// Check if target is in the left sorted part
	if pivot > 0 &&target >= nums[0] && target <= nums[pivot-1] {
		return binary(nums, 0, pivot-1, target)
	}
	// Otherwise search the right part
	return binary(nums, pivot, len(nums)-1, target)
}
