package particle

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Particle) Draw() {
	rl.DrawRectangle(int32(math.Round(p.Pos.X)-4), int32(math.Round(p.Pos.Y))-4, 8, 8, rl.SkyBlue)
}
