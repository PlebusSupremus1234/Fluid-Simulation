package simulation

import (
	"math/rand"

	"github.com/PlebusSupremus1234/Fluid-Simulation/src/particle"
)

func (sim *Simulation) SpawnParticles(mouseX, mouseY int32) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			x := float64(int(mouseX)+j*15) + 3*rand.Float64()
			y := float64(int(mouseY)+i*15) + 3*rand.Float64()

			badX := (x+sim.EPS > sim.VIEW_WIDTH) || (x-sim.EPS < 0)
			badY := (y+sim.EPS > sim.VIEW_HEIGHT) || (y-sim.EPS < 0)

			if !badX && !badY {
				sim.particles = append(sim.particles, particle.New(
					x, y,
					len(sim.particles),
				))
			}
		}
	}
}
