package vector

import "math"

func (v *Vector) Assign(vector Vector) {
	v.X = vector.X
	v.Y = vector.Y
}

func (v *Vector) Add(vector Vector, change bool) *Vector {
	x := v.X + vector.X
	y := v.Y + vector.Y

	if change {
		v.X = x
		v.Y = y
	}

	return New(x, y)
}

func (v *Vector) Subtract(vector Vector, change bool) *Vector {
	x := v.X - vector.X
	y := v.Y - vector.Y

	if change {
		v.X = x
		v.Y = y
	}

	return New(x, y)
}

func (v *Vector) Multiply(length float64, change bool) *Vector {
	x := v.X * length
	y := v.Y * length

	if change {
		v.X = x
		v.Y = y
	}

	return New(x, y)
}

func (v *Vector) Divide(length float64, change bool) *Vector {
	x := v.X / length
	y := v.Y / length

	if change {
		v.X = x
		v.Y = y
	}

	return New(x, y)
}

func (v *Vector) Dot(vector Vector) float64 {
	return v.X*vector.X + v.Y*vector.Y
}

func (v *Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() *Vector {
	length := v.Mag()

	if length == 0 {
		return New(1, 0)
	} else {
		return v.Divide(length, false)
	}
}
