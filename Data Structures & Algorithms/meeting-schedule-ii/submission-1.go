/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	mp := make(map[int]int)
	for _,i:= range intervals{
		mp[i.start]++
		mp[i.end]--
	}
	keys := make([]int, 0,len(mp))
	for k := range mp{
		keys = append(keys,k)
	}
	sort.Ints(keys)
	prev, res := 0,0
	for _,k := range keys{
		prev +=mp[k]
		if prev > res{
			res = prev
		}
	}
	return res
}
