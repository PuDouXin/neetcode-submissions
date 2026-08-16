func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i, num := range nums{
		if num > 0{
			break
		}
		if i>0 && num ==nums[i-1]{
			continue
		}
		
		j :=i+1
		k := len(nums)-1
		for j < k {
			tmp := nums[j]+nums[k]+num
			if tmp ==0{
				res = append(res, []int{num,nums[j],nums[k]})
				j++
				i--
				for j<k && nums[j]==nums[j-1]{
					j++
				}
			}else if tmp>0{
					k--
			}else{
				j++
			}
		}
	}
	return res
}
