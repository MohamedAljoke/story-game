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
	Loop     *application.GameLoop
}

func NewEngine(world *domain.World, loop *application.GameLoop) *Game {
	return &Game{
		World:    world,
		Input:    &EbitenInput{},
		Renderer: &EbitenRenderer{},
		Loop:     loop,
	}
}

func (g *Game) Update() error {
	if g.Loop.ProcessCommands(g.Input.ReadCommands()) {
		return eb.Termination
	}
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	g.Renderer.Draw(screen, g.World)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return domain.WorldWidth, domain.WorldHeight
}

func (g *Game) Run() error {
	eb.SetWindowSize(domain.WorldWidth, domain.WorldHeight)
	eb.SetWindowTitle("Dream Walker")
	eb.SetWindowResizingMode(eb.WindowResizingModeEnabled)

	err := eb.RunGame(g)
	if err == eb.Termination {
		return nil
	}
	return err
}
