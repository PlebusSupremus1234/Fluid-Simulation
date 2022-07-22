package vector

type Vector struct {
	X float64
	Y float64
}

func New(x, y float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}
