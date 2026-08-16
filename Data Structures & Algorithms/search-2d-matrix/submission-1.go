func binary_search(nums []int, target int, l int, r int) bool {
	if l >= r {
		if target == nums[l] {
			return true
		} else {
			return false
		}
	}

	m := (l + r) / 2
	if target == nums[m] {
		return true
	} else if target < nums[m] {
		return binary_search(nums, target, l, m)
	} else {
		return binary_search(nums, target, m+1, r)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	for _,m:= range matrix{
		if target <= m[len(m)-1] &&target>=m[0]{
			return binary_search(m,target,0,len(m)-1)
		}
	}
	return false
}
