func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand)%groupSize >0{
		return false
	}
	freq := make(map[int]int)
	for i := range hand{
		freq[hand[i]]++
	}
	for _,num := range hand{
		start := num
		for freq[start-1]>0{
			start --
		}
		for  start <=num{
			for freq[start]>0{
				for i := start; i<start+groupSize;i++{
					if freq[i] ==0{
						return false
					}
					freq[i]--
				}
			}
			start++
		}
	}
	return true
}
