package raylib

import (
	"fmt"

	"story-game/internal/domain"
)

type RaylibRenderer struct{}

func (r *RaylibRenderer) Draw(screen any, world *domain.World) {
	for id, pos := range world.Positions {
		name := world.Characters[id].Name
		fmt.Printf("  %s: (%.0f, %.0f)\n", name, pos.X, pos.Y)
	}
}
