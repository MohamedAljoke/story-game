package main

import (
	"log"

	"story-game/internal/adapter/raylib"
	"story-game/internal/application"
	"story-game/internal/domain"
)

func main() {
	world := domain.NewWorld()

	loop := &application.GameLoop{
		MoveChar: &application.MoveCharacter{World: world},
	}

	input := raylib.NewRaylibInput()

	game := &raylib.Game{
		World:    world,
		Input:    input,
		Renderer: &raylib.RaylibRenderer{},
		Loop:     loop,
	}

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
