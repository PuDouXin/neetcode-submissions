func sortColors(nums []int) {
    counts := make([]int, 3)
	for _, nums := range nums{
		counts[nums]++
	}
	index :=0
	for i:=0; i<3; i++{
		for counts[i] > 0{
			counts[i]--
			nums[index]=i
			index++
		}
	}
}
