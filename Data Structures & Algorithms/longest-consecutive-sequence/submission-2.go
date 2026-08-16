
func longestConsecutive(nums []int) int {
	if len(nums) ==0 {
		return 0
	}
 	sort.Ints(nums)
	res:=0
	final := 0
	s:=nums[0]
	i:=0
	for i<len(nums){
		if s !=nums[i]{
			s = nums[i]
			res =0
		}
		for i<len(nums)&& s == nums[i]{
			i++
		}
		res++
		s++
		if res > final{
			final = res
		}
	}
	return final
}
