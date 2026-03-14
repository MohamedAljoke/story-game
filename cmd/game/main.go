package main

import (
	"log"

	"story-game/internal/application"
	"story-game/internal/domain"
	ebitenpkg "story-game/internal/ebiten"

	eb "github.com/hajimehoshi/ebiten/v2"
)

func main() {
	world := domain.NewWorld()

	game := &ebitenpkg.Game{
		World:    world,
		Input:    &ebitenpkg.EbitenInput{},
		Renderer: &ebitenpkg.EbitenRenderer{},
		MoveChar: &application.MoveCharacter{World: world},
	}

	eb.SetWindowSize(domain.WorldWidth, domain.WorldHeight)
	eb.SetWindowTitle("Story Game")
	eb.SetWindowResizingMode(eb.WindowResizingModeEnabled)

	if err := eb.RunGame(game); err != nil && err != eb.Termination {
		log.Fatal(err)
	}
}
