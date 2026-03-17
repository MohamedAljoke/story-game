package main

import (
	"log"

	ebitenpkg "story-game/internal/adapter/ebiten"
	"story-game/internal/application"
	"story-game/internal/domain"
	"story-game/internal/loader"
)

func main() {
	tm, err := loader.LoadTiledMap("assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	tileset, err := ebitenpkg.NewTileset("assets/maps/images/TilesetFloor.png", tm.TileSize, tm.TileSize)
	if err != nil {
		log.Fatal(err)
	}

	world := domain.NewWorld(tm)

	loop := &application.GameLoop{
		MoveChar: &application.MoveCharacter{World: world},
		World:    world,
	}

	game := ebitenpkg.NewEngine(world, loop, tileset)

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
