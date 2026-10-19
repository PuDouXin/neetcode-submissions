type CountSquares struct {
    pts []Point
	ptsCount map[Point]int
}

type Point struct{
	x, y int
}

func Constructor() CountSquares {
    return CountSquares{
		pts: make([]Point, 0),
		ptsCount: make(map[Point]int),
	}
}

func (this *CountSquares) Add(point []int)  {
    p := Point{point[0], point[1]}
	this.pts = append(this.pts, p)
	this.ptsCount[p]++
}

func (this *CountSquares) Count(point []int) int {
    res := 0
	px, py := point[0], point[1]
	for _, pt := range this.pts {
		if abs(pt.x - px)!=abs(pt.y-py) || pt.x == px || pt.y == py{
			continue
		}
		p1 := Point{px,pt.y}
		p2 := Point{pt.x, py}

		res+=this.ptsCount[p1]*this.ptsCount[p2]
	}
	return res
}

func abs(x int)int{
	if x<0{
		return -x
	}
	return x
}
