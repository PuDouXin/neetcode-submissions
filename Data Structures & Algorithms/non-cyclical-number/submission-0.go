func isHappy(n int) bool {
    vis := make(map[int]bool)
	

	var res int

	for n >0{
		m := n%10
		res += m*m
		n =n/10
	}
	if res ==1{
		return true
	}

	t := res
	for !vis[t]{
		var tmp int
		vis[t] = true

		for t >0{
			m := t%10
			tmp += m*m
			t =t/10
		}
		if tmp ==1{
			return true
		}
		t = tmp
	}
	return false
}
