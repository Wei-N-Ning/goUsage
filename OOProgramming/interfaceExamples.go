package OOProgramming

import "math"

type PointI interface {
	distanceTo(PointI) float64
}

type VectorI interface {
	length() float64
	dotProduct(VectorI) float64
	crossProduct(VectorI) VectorI
}

type Point struct {
	x float64
	y float64
	z float64
}

func (p Point) distanceTo(other Point) float64 {
	x := other.x - p.x
	y := other.y - p.y
	z := other.z - p.z
	return math.Sqrt(x * x + y * y + z * z)
}
