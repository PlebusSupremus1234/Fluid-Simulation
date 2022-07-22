package main

import (
	sim "github.com/PlebusSupremus1234/Fluid-Simulation/src/simulation"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var width, height float64 = 1536, 718

	simulation := sim.New(width, height)
	settings := sim.Settings{
		SpaceDown: false,
	}

	rl.InitWindow(int32(width), int32(height), "Fluid Simulation")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		if rl.IsKeyDown(rl.KeySpace) {
			if !settings.SpaceDown {
				mouseX := rl.GetMouseX()
				mouseY := rl.GetMouseY()

				simulation.SpawnParticles(mouseX, mouseY)

				settings.SpaceDown = true
			}
		} else {
			settings.SpaceDown = false
		}

		simulation.Run()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
