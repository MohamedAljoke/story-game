package main

import (
	"log"

	ebitenpkg "story-game/internal/adapter/ebiten"
	"story-game/internal/application"
	"story-game/internal/domain"
	"story-game/internal/loader"
)

func main() {
	tm, tilesets, err := loader.LoadTiledMap("assets/maps/spawn.json", 2)
	if err != nil {
		log.Fatal(err)
	}

	tiles, err := ebitenpkg.NewTileRenderer(tilesets, 2.0)
	if err != nil {
		log.Fatal(err)
	}

	world := domain.NewWorld(tm)

	loop := &application.GameLoop{
		MoveChar: &application.MoveCharacter{World: world},
		World:    world,
	}

	game := ebitenpkg.NewEngine(world, loop, tiles)

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
