package block

import (
	"github.com/PlebusSupremus1234/Fluid-Simulation/src/vector"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Pos vector.Vector
	W   int32
	H   int32
}

func New(x, y float64, w, h int32) *Block {
	return &Block{
		Pos: *vector.New(x, y),
		W:   w,
		H:   h,
	}
}

func (b *Block) Draw() {
	rl.DrawRectangle(int32(b.Pos.X), int32(b.Pos.Y), b.W, b.H, rl.Brown)
}
