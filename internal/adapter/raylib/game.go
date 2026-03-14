package raylib

import (
	"story-game/internal/application"
	"story-game/internal/domain"
	"story-game/internal/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	World    *domain.World
	Input    engine.InputReader
	Renderer engine.Renderer
	Loop     *application.GameLoop
}

func (g *Game) Run() error {
	rl.InitWindow(domain.WorldWidth, domain.WorldHeight, "Dream Walker")
	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyNull)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		if g.Loop.ProcessCommands(g.Input.ReadCommands()) {
			return nil
		}
		rl.BeginDrawing()
		g.Renderer.Draw(nil, g.World)
		rl.EndDrawing()
	}
	return nil
}
