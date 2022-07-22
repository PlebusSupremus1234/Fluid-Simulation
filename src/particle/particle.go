package particle

import "github.com/PlebusSupremus1234/Fluid-Simulation/src/vector"

type Particle struct {
	Pos    vector.Vector // Position
	OldPos vector.Vector // Old Position
	Vel    vector.Vector // Velocity
	Force  vector.Vector // Force

	Rho float64 // Density
	P   float64 // Pressure
	M   float64 // Mass

	Index int // Grid Index
}

func New(x, y float64, index int) *Particle {
	return &Particle{
		Pos:    *vector.New(x, y),
		OldPos: *vector.New(x, y),
		Vel:    *vector.New(0, 0),
		Force:  *vector.New(0, 0),

		Rho: 0,
		P:   0,
		M:   2.5,

		Index: index,
	}
}
