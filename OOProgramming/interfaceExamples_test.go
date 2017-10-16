package OOProgramming

import "testing"

// TODO: use table based test case
func TestPointExpectDistance(t *testing.T) {
	p1 := Point{1.0, 0, 0}
	p2 := Point{-1.0, 0, 0}
	dist := p1.distanceTo(p2)
	if 2.0 != dist {
		t.Error(dist)
	}
}


func TestExpectNotRun(t *testing.T) {
	type BadPoint struct {
		x int
		y int
		z int
	}
	//interface violation (IDE warning)
	//badPoint := BadPoint{}
	//p1 := Point{1.0, 0, 0}
}