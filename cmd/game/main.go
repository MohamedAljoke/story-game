package main

import (
	"log"

	ebitenpkg "story-game/internal/adapter/ebiten"
	"story-game/internal/application"
	"story-game/internal/domain"
)

func main() {
	world := domain.NewWorld()

	loop := &application.GameLoop{
		MoveChar: &application.MoveCharacter{World: world},
	}

	game := ebitenpkg.NewEngine(world, loop)

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
