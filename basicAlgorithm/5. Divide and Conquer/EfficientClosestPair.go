package main

// 这道题是自己COPY自己很久以前写的代码  这个写起来太累了吧

import (
	math "math"
	"sort"
)

type Pair struct {
	x, y float64
}

//模仿GOPL设计
type pairList struct {
	t    []*Pair
	less func(x, y *Pair) bool
}

func (p pairList) Len() int {
	return len(p.t)
}

func (p pairList) Less(i, j int) bool {
	return p.less(p.t[i], p.t[j])
}

func (p pairList) Swap(i, j int) {
	p.t[i], p.t[j] = p.t[j], p.t[i]
}

func byX(i, j *Pair) bool {
	return i.x < j.x
}

func byY(i, j *Pair) bool {
	return i.y < j.y
}

func EfficientClosestPair(pairs []Pair) ([]*Pair, float64) {
	length := len(pairs)
	temp := make([]*Pair, length)
	temp2 := make([]*Pair, length)
	for i := 0; i < len(pairs); i++ {
		temp[i] = &pairs[i]
		temp2[i] = &pairs[i]
	}
	P := pairList{t: temp, less: byX}
	Q := pairList{t: temp2, less: byY}
	sort.Sort(P)
	sort.Sort(Q)
	return efficientClosestPair(&P, &Q, 0, len(pairs)-1)
}

//P X排序 Q Y排序    意为从 low->high 中找到最短距离
func efficientClosestPair(P, Q *pairList, low, high int) ([]*Pair, float64) {
	if high-low+1 <= 3 {
		return ViolenceClosestPair(P.t[low:high])
	}
	var result []*Pair
	mid := (low + high) >> 1
	x0 := P.t[mid].x
	result1, distantce1 := efficientClosestPair(P, Q, low, mid)
	result2, distantce2 := efficientClosestPair(P, Q, mid+1, high)
	var minDistance float64
	if distantce1 < distantce2 {
		minDistance = distantce1
		result = result1
	} else {
		minDistance = distantce2
		result = result2
	}
	middle := []*Pair{}
	for i := 0; i < Q.Len(); i++ {
		if math.Abs(Q.t[i].x-x0) <= minDistance {
			middle = append(middle, Q.t[i])
		}
	}
	for i := 0; i < len(middle); i++ {
		for j := i + 1; j < i+7 && j < len(middle); j++ {
			distance := math.Sqrt((middle[i].x-middle[j].x)*(middle[i].x-middle[j].x) + (middle[i].y-middle[j].y)*(middle[i].y-middle[j].y))
			if minDistance > distance {
				minDistance = distance
				result = []*Pair{middle[i], middle[j]}
			}
		}
	}
	return result, minDistance
}

func ViolenceClosestPair(pairs []*Pair) (result []*Pair, minDistance float64) {
	minDistance = math.MaxFloat64
	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			distance := math.Sqrt((pairs[i].x-pairs[j].x)*(pairs[i].x-pairs[j].x) + (pairs[i].y-pairs[j].y)*(pairs[i].y-pairs[j].y))
			if minDistance > distance {
				minDistance = distance
				result = []*Pair{pairs[i], pairs[j]}
			}
		}
	}
	return
}
