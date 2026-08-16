
func minEatingSpeed(piles []int, h int) int {
	maximum := 0
	for _, p:= range piles{
		if p>maximum{
			maximum = p
		}
	}
	res:= maximum
	l,r := 1, maximum

	for l<=r{
		m :=(l+r)/2
		total := 0
		for _,p := range piles{
			total += int(math.Ceil(float64(p)/float64(m)))
		}
		if total<=h{
			res = m
			r = m-1
		}else{
			l=m+1
		}
	}

	return res
	
}
