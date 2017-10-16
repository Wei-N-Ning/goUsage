package OOProgramming

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
	return 0.0
}
