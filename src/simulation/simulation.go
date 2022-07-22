package simulation

import (
	"math"
	"math/rand"

	"github.com/PlebusSupremus1234/Fluid-Simulation/src/block"
	"github.com/PlebusSupremus1234/Fluid-Simulation/src/particle"
	"github.com/PlebusSupremus1234/Fluid-Simulation/src/vector"
)

type Simulation struct {
	particles  []*particle.Particle     // Simulation particles
	neighbours [][]*particle.Particle   // Neighbours for each particle
	grid       [][][]*particle.Particle // Grid for faster neighbour lookup

	blocks []*block.Block

	H   float64 // Radius
	HSQ float64 // Radius^2

	REST_DENSITY float64 // Rest density
	GAS_CONSTANT float64 // Gas constant

	VISCOSITY float64 // Viscosity

	SURFACE_TENSION float64 // Surface tension

	GRAVITY vector.Vector // Gravity
	DT      float64       // Integration timestep

	// Smoothing kernels
	POLY6      float64
	SPIKY_GRAD float64
	VISC_LAP   float64
	COHESION   float64

	EPS           float64 // Boundary epsilon
	BOUND_DAMPING float64 // Boundary damping

	VIEW_WIDTH  float64 // View width
	VIEW_HEIGHT float64 // View height

	COLS float64 // Columns
	ROWS float64 // Rows
}

func New(width, height float64) *Simulation {
	var particles []*particle.Particle

	for i := 0; i < 30; i++ {
		for j := 0; j < 20; j++ {
			particles = append(particles, particle.New(
				float64(100+j*15)+3*rand.Float64(),
				float64(200+i*15)+3*rand.Float64(),

				len(particles), // Index
			))
		}
	}

	var H float64 = 16

	return &Simulation{
		particles:  particles,
		neighbours: [][]*particle.Particle{},
		grid:       [][][]*particle.Particle{},

		blocks: []*block.Block{block.New(50, 50, 50, 50)},

		H:   H,
		HSQ: H * H,

		REST_DENSITY: 300,
		GAS_CONSTANT: 2000,

		VISCOSITY: 200,

		SURFACE_TENSION: 0.001,

		GRAVITY: *vector.New(0, 10),
		DT:      0.0007,

		POLY6:      315 / (64 * math.Pi * math.Pow(H, 8)),
		SPIKY_GRAD: -15 / (math.Pi * math.Pow(H, 5)),
		VISC_LAP:   40 / (math.Pi * math.Pow(H, 5)),
		COHESION:   32 / (math.Pi * math.Pow(H, 9)),

		EPS:           H,
		BOUND_DAMPING: -0.5,

		VIEW_WIDTH:  width,
		VIEW_HEIGHT: height,

		COLS: math.Ceil(height / H),
		ROWS: math.Ceil(width / H),
	}
}

func (sim *Simulation) Draw() {
	for _, i := range sim.particles {
		i.Draw()
	}

	for _, i := range sim.blocks {
		i.Draw()
	}

	// for i := range sim.grid {
	// 	for j := range sim.grid[i] {
	// 		if len(sim.grid[i][j]) != 0 {
	// 			rl.DrawRectangle(int32(j)*int32(sim.H), int32(i)*int32(sim.H), int32(sim.H), int32(sim.H), rl.Beige)
	// 		}
	// 	}
	// }
}

func (sim *Simulation) Run() {
	sim.UpdateGrid()
	sim.UpdateNeighbours()

	sim.computeDensityPressure()
	sim.computeForces()
	sim.integrate()

	sim.Draw()
}
