package simulation

import (
	"math"

	"github.com/PlebusSupremus1234/Fluid-Simulation/src/vector"
)

func (sim *Simulation) computeDensityPressure() {
	for i, Pi := range sim.particles {
		Pi.Rho = 0

		for _, Pj := range sim.neighbours[i] {
			rij := Pj.Pos.Subtract(Pi.Pos, false)
			r2 := rij.Dot(*rij)

			if rij.Mag() < sim.H {
				Pi.Rho += Pj.M * sim.POLY6 * math.Pow(sim.HSQ-r2, 3)
			}
		}

		Pi.Rho += Pi.M * sim.POLY6 * math.Pow(sim.HSQ, 3)

		Pi.P = sim.GAS_CONSTANT * (Pi.Rho - sim.REST_DENSITY)
	}
}

func (sim *Simulation) computeForces() {
	for i, Pi := range sim.particles {
		// Pressure force
		// In LaTeX: F^{pressure}_i = -\sum_j m_j \frac{p_i + p_j}{2\rho_j} \nabla W(r_i - r_j, h)
		Fpressure := vector.New(0, 0)

		// Viscosity force
		// In LaTeX: F^{viscosity}_i = \mu \sum_j m_j \frac{v_j - v_i}{\rho_j} \nabla^2 W(r_i - r_j, h)
		Fviscosity := vector.New(0, 0)

		// Cohesion force
		// In LaTeX: F^{cohesion}_{i \leftarrow j} = -\alpha m_i m_j \frac{x_i - x_j}{||x_i - x_j||} W^{cohesion}(||x_i - x_j||)
		Fcohesion := vector.New(0, 0)

		for _, Pj := range sim.neighbours[i] {
			rij := Pj.Pos.Subtract(Pi.Pos, false)
			r := rij.Mag()

			if r < sim.H {
				normalized := rij.Normalize()

				// Compute pressure force
				Wp := sim.SPIKY_GRAD * math.Pow(sim.H-r, 3)
				multiplierp := Pj.M * (Pi.P + Pj.P) / (2 * Pj.Rho) * Wp
				FpressureIncrease := normalized.Multiply(multiplierp, false)

				Fpressure.Add(*FpressureIncrease, true)

				// Compute viscosity force
				velDifference := Pj.Vel.Subtract(Pi.Vel, false)

				Wv := sim.VISC_LAP * (sim.H - r)
				multiplierv := Pj.M / Pj.Rho * Wv
				FviscosityIncrease := velDifference.Multiply(multiplierv, false)

				Fviscosity.Add(*FviscosityIncrease, true)

				// Compute cohesion force
				Wc := sim.Wcohesion(r)
				multiplierc := -sim.SURFACE_TENSION * Pi.M * Pj.M * Wc
				FcohesionIncrease := normalized.Multiply(multiplierc, false)

				Fcohesion.Add(*FcohesionIncrease, true)
			}
		}

		Fpressure.Multiply(-1, true)
		Fviscosity.Multiply(sim.VISCOSITY, true)
		Fgravity := sim.GRAVITY.Multiply(Pi.M/Pi.Rho, false)

		sum := Fcohesion.Add(*Fpressure.Add(*Fviscosity, false).Add(*Fgravity, false), false)
		Pi.Force.Assign(*sum)
	}
}

func (sim *Simulation) Wcohesion(r float64) float64 {
	a := math.Pow(sim.H-r, 3)
	b := math.Pow(r, 3)

	if 2*r > sim.H && r <= sim.H {
		return sim.COHESION * a * b
	} else if r > 0 && 2*r <= sim.H {
		return sim.COHESION*2*a*b - math.Pow(sim.HSQ, 3)/64
	} else {
		return 0
	}
}

func (sim *Simulation) integrate() {
	for _, P := range sim.particles {
		if P.Pos.X != P.OldPos.X {
			P.OldPos.X = P.Pos.X
		}
		if P.Pos.Y != P.OldPos.Y {
			P.OldPos.Y = P.Pos.Y
		}

		P.Vel.Add(*P.Force.Multiply(sim.DT/P.Rho, false), true)
		P.Pos.Add(*P.Vel.Multiply(sim.DT, false), true)

		if P.Pos.X-sim.EPS < 0 {
			P.Vel.X *= sim.BOUND_DAMPING
			P.Pos.X = sim.EPS
		}
		if P.Pos.X+sim.EPS > sim.VIEW_WIDTH {
			P.Vel.X *= sim.BOUND_DAMPING
			P.Pos.X = sim.VIEW_WIDTH - sim.EPS
		}
		if P.Pos.Y-sim.EPS < 0 {
			P.Vel.Y *= sim.BOUND_DAMPING
			P.Pos.Y = sim.EPS
		}
		if P.Pos.Y+sim.EPS > sim.VIEW_HEIGHT {
			P.Vel.Y *= sim.BOUND_DAMPING
			P.Pos.Y = sim.VIEW_HEIGHT - sim.EPS
		}
	}
}
