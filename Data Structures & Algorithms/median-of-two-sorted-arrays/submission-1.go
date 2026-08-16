func binary(nums []int,v int)int{
	l,r :=0, len(nums)-1
	for l<=r{
		m := (l+r)/2
		if v>nums[m]{
			l=m+1
		}else{
			r=m-1
		}
	}
	return l
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//l,r := 0,0
	merged := nums1
	for _,v:= range nums2{
		idx := binary(merged,v)
		merged = append(merged,0)
		copy(merged[idx+1:],merged[idx:])
		merged[idx]=v

	}
	m:=(len(nums1)+len(nums2))/2
	if (len(nums1)+len(nums2))%2==0{
		
		return (float64(merged[m])+float64(merged[m-1]))/2
	}
	return float64(merged[m])
}
