func reverseBits(n int) int {
	//var res uint32 = 0
	res :=0
	for i := 0; i<32; i++{
		bit := (n>>i)&1
		if bit >0{
			res |= (bit <<(31-i))
		}
	}
	return res
}
 