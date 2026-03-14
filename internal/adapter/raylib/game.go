package raylib

import (
	"fmt"

	"story-game/internal/application"
	"story-game/internal/domain"
	"story-game/internal/engine"
)

type Game struct {
	World    *domain.World
	Input    engine.InputReader
	Renderer engine.Renderer
	Loop     *application.GameLoop
}

func (g *Game) Run() error {
	fmt.Println("Dream Walker (terminal stub)")
	fmt.Println("w/a/s/d to move, q to quit")

	for {
		fmt.Print("> ")
		cmds := g.Input.ReadCommands()
		if g.Loop.ProcessCommands(cmds) {
			fmt.Println("Goodbye!")
			return nil
		}
		g.Renderer.Draw(nil, g.World)
	}
}
