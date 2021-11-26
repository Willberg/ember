package test

import (
	"fmt"
	"testing"
)

type Point struct {
	x, y int64
}

func Distance(a, b *Point) (d int64) {
	d = a.x - b.x
	return
}

func (p Point) Distance(o *Point) int64 {
	p.x = 9
	return p.x - o.x
}

func (p Point) OriginalDistance(o *Point) int64 {
	return p.x - o.x
}

func (p *Point) PointerDistance(o *Point) (d int64) {
	d = p.x - o.x
	return
}

func (p *Point) ChangeX() {
	p.x = 8
}

func TestDistance(t *testing.T) {
	a, b := &Point{10, 9}, &Point{5, 5}
	fmt.Println(Distance(a, b))
	fmt.Println(a.Distance(b))
	fmt.Println(a.PointerDistance(b))
	fmt.Println(Point.Distance(*a, b))
	fmt.Println((*Point).PointerDistance(a, b))
	a.ChangeX()
	fmt.Println(a.x)
	fmt.Println(b.x)
	fmt.Println(a.OriginalDistance(b))
}
