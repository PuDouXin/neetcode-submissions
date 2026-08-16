func quickSort(nums []int){
	if len(nums)<2{
		return
	}

	left := 0
	right := len(nums)-1

	pivot := nums[right]
	i := left
	for j:=left;j<right;j++{
		if nums[j] < pivot {
			nums[i],nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]

	quickSort(nums[:i])
	quickSort(nums[i+1:])

}
func sortColors(nums []int) {
    quickSort(nums)
}
