

type TimeMap struct {
	maps map[string]map[int]string

}

func Constructor() TimeMap {
	maps:= make(map[string]map[int]string)
	return TimeMap{maps:maps}

}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _,exist := this.maps[key]; !exist{
		this.maps[key]=make(map[int]string)
	}
	this.maps[key][timestamp]=value
	}

func (this *TimeMap) Get(key string, timestamp int) string {
	 v,exist:=this.maps[key]
	 if !exist{
		return ""
	}
	seen:= 0
	for time,_:= range v{
		if time<=timestamp{
			if seen<time{
				seen=time
			}
		}
	}
	if seen==0{
		return ""
	}
	return v[seen]

}
