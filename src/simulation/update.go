package simulation

import (
	"math"

	"github.com/PlebusSupremus1234/Fluid-Simulation/src/particle"
)

func (sim *Simulation) UpdateNeighbours() {
	var newNeighbours [][]*particle.Particle

	for range sim.particles {
		newNeighbours = append(newNeighbours, []*particle.Particle{})
	}

	for i, Pi := range sim.particles {
		neighbours := sim.FindGridNeighbours(Pi)

		for _, Pj := range neighbours {
			rij := Pi.Pos.Subtract(Pj.Pos, false)
			r := rij.Mag()

			if r < sim.H {
				newNeighbours[i] = append(newNeighbours[i], Pj)
			}
		}
	}

	sim.neighbours = newNeighbours
}

func (sim *Simulation) UpdateGrid() {
	var newGrid [][][]*particle.Particle

	for i := 0; float64(i) < sim.COLS; i++ {
		var push [][]*particle.Particle

		for j := 0; float64(j) < sim.ROWS; j++ {
			push = append(push, []*particle.Particle{})
		}

		newGrid = append(newGrid, push)
	}

	for _, Pi := range sim.particles {
		x := int(math.Floor(Pi.Pos.X / sim.H))
		y := int(math.Floor(Pi.Pos.Y / sim.H))

		newGrid[y][x] = append(newGrid[y][x], Pi)
	}

	sim.grid = newGrid
}

func (sim *Simulation) FindGridNeighbours(P *particle.Particle) []*particle.Particle {
	var indexX int = int(math.Floor(P.Pos.X / sim.H))
	var indexY int = int(math.Floor(P.Pos.Y / sim.H))

	lessX := indexX-1 >= 0
	lessY := indexY-1 >= 0
	moreX := float64(indexX)+1 < sim.ROWS
	moreY := float64(indexY)+1 < sim.ROWS

	var found []*particle.Particle

	if lessX && lessY {
		found = append(found, sim.grid[indexY-1][indexX-1]...)
	}
	if lessX {
		found = append(found, sim.grid[indexY][indexX-1]...)
	}
	if lessX && moreY {
		found = append(found, sim.grid[indexY+1][indexX-1]...)
	}

	if lessY {
		found = append(found, sim.grid[indexY-1][indexX]...)
	}
	if moreY {
		found = append(found, sim.grid[indexY+1][indexX]...)
	}

	if moreX && lessY {
		found = append(found, sim.grid[indexY-1][indexX+1]...)
	}
	if moreX {
		found = append(found, sim.grid[indexY][indexX+1]...)
	}
	if moreX && moreY {
		found = append(found, sim.grid[indexY+1][indexX+1]...)
	}

	for _, Pi := range sim.grid[indexY][indexX] {
		if Pi != P {
			found = append(found, Pi)
		}
	}

	return found
}
