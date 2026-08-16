type tweet struct {
	id int
	c  int
}
type maxHeap struct {
	data []tweet
}

func Newheap() maxHeap {
	return maxHeap{data: []tweet{}}
}

func (m *maxHeap) Push(t tweet) {
	m.data = append(m.data, t)
	m.SiftUp(len(m.data) - 1)
}

func (m *maxHeap) Pop() tweet {
	if len(m.data) == 0 {
		return tweet{}
	}
	top := m.data[0]
	last := len(m.data) - 1
	m.data[0] = m.data[last]
	m.data = m.data[:last]
	m.SiftDown(0)
	return top
}

func (m *maxHeap) SiftDown(i int) {
	n := len(m.data)
	for {
		left := i*2 + 1
		right := i*2 + 2
		largest := i
		if left < n && m.data[left].c > m.data[largest].c {
			largest = left
		}
		if right < n && m.data[right].c > m.data[largest].c {
			largest = right
		}
		if largest == i {
			break
		}
		m.data[largest], m.data[i] = m.data[i], m.data[largest]
		i = largest
	}
}

func (m *maxHeap) SiftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if m.data[i].c <= m.data[parent].c {
			break
		}
		m.data[i], m.data[parent] = m.data[parent], m.data[i]
		i = parent
	}
}

type Twitter struct {
	tweetMap  map[int][]tweet //value is count,tweetId
	followMap map[int]map[int]bool
	count     int
}

func Constructor() Twitter {
	return Twitter{
		count:     0,
		followMap: make(map[int]map[int]bool),
		tweetMap:  make(map[int][]tweet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if this.tweetMap[userId] == nil {
		this.tweetMap[userId] = make([]tweet, 0)
	}
	this.tweetMap[userId] = append(this.tweetMap[userId], tweet{c: this.count, id: tweetId})
	this.count++
	l := len(this.tweetMap[userId])
	if l > 10 {
	
		this.tweetMap[userId] = this.tweetMap[userId][l-10:l]
	}
}

func (this *Twitter) getLastTweet(userId int) maxHeap {
	heap := Newheap()
	for follower := range this.followMap[userId] {
		tweets := this.tweetMap[follower]
		if len(tweets) > 0 {
			for i := len(tweets) - 1; i >= 0; i-- {
				t := tweets[i]
				heap.Push(t)
			}

		}
	}

	return heap
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0)
	if this.followMap[userId] == nil {
		this.followMap[userId] = make(map[int]bool)
	}
	this.followMap[userId][userId] = true
	heap := this.getLastTweet(userId)
	for len(res) < 10 && len(heap.data) > 0 {
		t := heap.Pop()
		res = append(res, t.id)
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.followMap[followerId] == nil {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followMap[followerId] != nil {
		delete(this.followMap[followerId], followeeId)
	}
}

