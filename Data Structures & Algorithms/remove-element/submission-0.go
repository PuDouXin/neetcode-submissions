func removeElement(nums []int, val int) int {
	res:=[]int{}
	
    for _,n := range nums{
		if n!=val{
			res = append(res,n)
		}
	}
	for i:=0;i<len(res);i++{
		nums[i]=res[i]
	}
	return len(res)
}
