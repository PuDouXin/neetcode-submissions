/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) ==0{
		return true
	}
	sort.Slice(intervals, func (a, b int)bool{
		return intervals[a].start<intervals[b].start
	})
	
	 maxV := intervals[0].end

	for i := 1; i<len(intervals);i++{
		if intervals[i].start < maxV{
			return false
		}
		maxV = intervals[i].end
	}
	return true
}
