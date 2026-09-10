func mergeTriplets(triplets [][]int, target []int) bool {
    good := make(map[int]bool)
	for i := range triplets{
		if triplets[i][0]>target[0] || triplets[i][1]>target[1] ||triplets[i][2]>target[2]{
			continue
		}
		for j,v := range triplets[i]{
			if v == target[j]{
				good[j] = true
			}
		}
	}
	return len(good)==3
}
