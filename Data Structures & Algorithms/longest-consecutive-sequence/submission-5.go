
func longestConsecutive(nums []int) int {
	if len(nums)==0{
		return 0
	}
 	sort.Ints(nums)
	res:=1
	final := 1
	s:=nums[0]
	i:=1
	for i<len(nums){
		for i<len(nums)&& s == nums[i]{
			i++
		}
		if i<len(nums){
		if nums[i]-s!=1{
			res = 1
		}else{
			res++
		}
		s = nums[i]
		i++
		}
		if res > final{
			final = res
		}
	}
	return final
}
