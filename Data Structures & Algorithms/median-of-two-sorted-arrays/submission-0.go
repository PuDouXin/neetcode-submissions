func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l,r := 0,0
	merged := make([]int,0,len(nums1)+len(nums2))
	for l < len(nums1) && r < len(nums2){
		if nums1[l]<nums2[r]{
			merged = append(merged,nums1[l])
			l++
		}else{
			merged = append(merged,nums2[r])
			r++
		}
	}
	for l<len(nums1){
		merged = append(merged,nums1[l])
		l++
	}
	for r<len(nums2){
		merged = append(merged,nums2[r])
			r++
	}

	m:=(len(nums1)+len(nums2))/2
	if (len(nums1)+len(nums2))%2==0{
		
		return (float64(merged[m])+float64(merged[m-1]))/2
	}
	return float64(merged[m])
}
