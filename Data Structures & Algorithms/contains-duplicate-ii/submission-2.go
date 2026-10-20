func containsNearbyDuplicate(nums []int, k int) bool {
	
	i,j := 0, 1
	for j <= len(nums)-1{
		if nums[i] == nums[j]{
			return true
		}
		j++
		if j-i >k{
			i++
			j = i + 1
		}
	}
	return false
}
