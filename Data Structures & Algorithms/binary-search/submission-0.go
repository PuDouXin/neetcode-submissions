func binary_search(nums[] int, target int, l int, r int)int{
   if l>=r{
	if target ==nums[l]{
		return l
	}else{
		return -1
	}
   }

	m := (l+r)/2
	if target == nums[m]{
		return m
	}else if target<nums[m]{
		return binary_search(nums,target,l,m)
	}else{
		return binary_search(nums,target,m+1,r)
	}
}
func search(nums []int, target int) int {
	return binary_search(nums,target,0,len(nums)-1)
}
