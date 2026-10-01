func reverse(x int) int {
	MIN := -(1<<31)
	MAX := (1<<31)-1
	res :=0
	for x!=0{
		d := int(math.Mod(float64(x),10))
		x = x/10

		if res > MAX/10 || (res ==MAX/10 && d > MAX%10){
			return 0
		}
		if res < MIN/10 || (res ==MIN/10 && d < MIN%10){
			return 0
		}
		res = (res *10)+d
		
	}

	return res
}
