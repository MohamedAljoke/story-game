package raylib

import (
	"story-game/internal/domain"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RaylibRenderer struct{}

func (r *RaylibRenderer) Draw(screen any, world *domain.World) {
	rl.ClearBackground(rl.Color{R: 0x1a, G: 0x1a, B: 0x2e, A: 0xff})

	gridColor := rl.Color{R: 0x2a, G: 0x2a, B: 0x3e, A: 0xff}
	for x := int32(0); x < domain.WorldWidth; x += domain.TileSize {
		rl.DrawLine(x, 0, x, domain.WorldHeight, gridColor)
	}
	for y := int32(0); y < domain.WorldHeight; y += domain.TileSize {
		rl.DrawLine(0, y, domain.WorldWidth, y, gridColor)
	}

	charColor := rl.Color{R: 0x00, G: 0xd4, B: 0xff, A: 0xff}
	for id, pos := range world.Positions {
		_ = world.Characters[id]
		rl.DrawRectangle(int32(pos.X), int32(pos.Y), domain.TileSize, domain.TileSize, charColor)
	}

	rl.DrawText("WASD / Arrow keys to move | ESC to quit", 10, 10, 10, rl.White)
}
