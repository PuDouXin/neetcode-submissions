func productExceptSelf(nums []int) []int {
	res := 	make([]int, len(nums))
	totalProduct := int(1)
	zeroCount :=0
	for _,n := range nums{
		if n==0{
			zeroCount ++
			continue
		}
		totalProduct *= n
	}
	for i,n := range nums {
		if zeroCount ==1 {
			if n == 0 {
				res[i] = totalProduct
			}else {
				res[i] = 0
			}
		}else if zeroCount == 0{
			res[i] = totalProduct/n
		}else{
			res[i] =0
		}
	}
	return res
}
