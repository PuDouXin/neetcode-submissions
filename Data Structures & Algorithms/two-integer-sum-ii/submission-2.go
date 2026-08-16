func twoSum(numbers []int, target int) []int {
	
	l,r := 0, len(numbers)-1
	for l<r{
		// for numbers[r]>target{
		// 		r--
		// }
		tmp := numbers[l]+numbers[r]
		if numbers[l]+numbers[r]==target{
			return []int{l+1,r+1}
		}else if tmp > target{
			r--
		}else{
			l++
		}
	}
	return []int{1,2}
}
