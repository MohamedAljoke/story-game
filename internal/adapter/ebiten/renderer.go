package ebiten

import (
	"image/color"

	"story-game/internal/domain"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type EbitenRenderer struct{}

func (r *EbitenRenderer) Draw(screen *ebiten.Image, world *domain.World) {

	// Background
	screen.Fill(hexColor(0x1a1a2e))

	// Grid
	gridColor := hexColor(0x2a2a3e)
	for x := 0; x < domain.WorldWidth; x += domain.TileSize {
		vector.StrokeLine(screen, float32(x), 0, float32(x), domain.WorldHeight, 1, gridColor, false)
	}
	for y := 0; y < domain.WorldHeight; y += domain.TileSize {
		vector.StrokeLine(screen, 0, float32(y), domain.WorldWidth, float32(y), 1, gridColor, false)
	}

	// Characters
	charColor := hexColor(0x00d4ff)
	for id, pos := range world.Positions {
		_ = world.Characters[id]
		vector.DrawFilledRect(screen, float32(pos.X), float32(pos.Y), domain.TileSize, domain.TileSize, charColor, false)
	}

	// HUD
	ebitenutil.DebugPrint(screen, "WASD / Arrow keys to move | ESC to quit")
}

func hexColor(v uint32) color.RGBA {
	return color.RGBA{R: uint8(v >> 16), G: uint8(v >> 8), B: uint8(v), A: 0xff}
}
