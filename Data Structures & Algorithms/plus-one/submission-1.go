func plusOne(digits []int) []int {
    var res []int
	n := len(digits)
	d := 1
	//digits[n-1]++
	for i := n-1; i>=0; i--{
		t:= digits[i]+d
		if t >= 10{
			d = 1
			t =t-10
		}else{
			d = 0
		}
		res = append([]int{t}, res...)
	}
	if d == 1{
		res = append([]int{1}, res...)
	}
	return res
	
}
