type pairs struct{
	timestamp int
	value string
}

type TimeMap struct {
	maps map[string][]pairs
}

func Constructor() TimeMap {
	maps:= make(map[string][]pairs)
	return TimeMap{maps:maps}

}


func (this *TimeMap) Set(key string, value string, timestamp int) {
	// if _,exist := this.maps[key]; !exist{
	// 	this.maps[key]=make([]pairs)
	// }
	this.maps[key]= append(this.maps[key], pairs{timestamp:timestamp, value:value})
	}

func (this *TimeMap) Get(key string, timestamp int) string {
	 v,exist:=this.maps[key]
	 if !exist{
		return ""
	}
	l:= 0
	r:= len(v)-1
	for l<=r{
		m:=(l+r)/2
		if timestamp>=v[m].timestamp{
			if m==len(v)-1 || v[m+1].timestamp>timestamp{
				return v[m].value
			}
			l=m+1
		}else{
			r=m-1
		}
	}
	return ""

}
