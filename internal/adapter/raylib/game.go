package raylib

import (
	"story-game/internal/application"
	"story-game/internal/domain"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	World    *domain.World
	Input    *RaylibInput
	Renderer *RaylibRenderer
	Loop     *application.GameLoop
}

func NewEngine(world *domain.World, loop *application.GameLoop) *Game {
	return &Game{
		World:    world,
		Input:    &RaylibInput{},
		Renderer: &RaylibRenderer{},
		Loop:     loop,
	}
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
