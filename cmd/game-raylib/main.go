package main

import (
	"log"

	raylibpkg "story-game/internal/adapter/raylib"
	"story-game/internal/application"
	"story-game/internal/domain"
)

func main() {
	world := domain.NewWorld()

	loop := &application.GameLoop{
		MoveChar: &application.MoveCharacter{World: world},
	}

	game := &raylibpkg.Game{
		World:    world,
		Input:    &raylibpkg.RaylibInput{},
		Renderer: &raylibpkg.RaylibRenderer{},
		Loop:     loop,
	}

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
