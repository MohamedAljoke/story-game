package ebiten

import (
	"story-game/internal/application"
	"story-game/internal/domain"
	"story-game/internal/engine"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	World    *domain.World
	Input    engine.InputReader
	Renderer engine.Renderer
	MoveChar *application.MoveCharacter
}

func (g *Game) Update() error {
	for _, cmd := range g.Input.ReadCommands() {
		switch cmd {
		case engine.CmdQuit:
			return eb.Termination
		case engine.CmdMoveUp:
			g.MoveChar.Execute("player", domain.DirUp)
		case engine.CmdMoveDown:
			g.MoveChar.Execute("player", domain.DirDown)
		case engine.CmdMoveLeft:
			g.MoveChar.Execute("player", domain.DirLeft)
		case engine.CmdMoveRight:
			g.MoveChar.Execute("player", domain.DirRight)
		}
	}
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	g.Renderer.Draw(screen, g.World)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return domain.WorldWidth, domain.WorldHeight
}
