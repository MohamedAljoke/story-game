package sprites

import (
	"story-game/internal/adapter/ebiten/helpers"
	"story-game/internal/domain"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type CatSprite struct {
	frames      [][]*eb.Image
	animFrame   int
	animCounter int
}

func NewCatSprite() *CatSprite {

	sheet, err := helpers.LoadSprite("assets/sprites/cat.png")
	if err != nil {
		panic(err)
	}

	frames := [][]*eb.Image{
		helpers.SliceRow(sheet, 1, 1, 32, 32, 3),   // DOWN
		helpers.SliceRow(sheet, 9, 13, 32, 32, 8),  // UP
		helpers.SliceRow(sheet, 3, 13, 32, 32, 4),  // LEFT
		helpers.SliceRow(sheet, 11, 13, 32, 32, 4), // RIGHT
	}

	return &CatSprite{
		frames: frames,
	}
}

func (c *CatSprite) Draw(screen *eb.Image, pos *domain.Position, dir domain.Direction) {

	c.animCounter++

	if c.animCounter >= 500 {
		c.animFrame = (c.animFrame + 1) % len(c.frames[0])
		c.animCounter = 0
	}

	row := c.getRow(dir)
	frame := c.frames[row][c.animFrame]

	w, h := frame.Bounds().Dx(), frame.Bounds().Dy()

	op := &eb.DrawImageOptions{}

	op.GeoM.Translate(
		pos.X+float64(domain.TileSize-w)/2,
		pos.Y+float64(domain.TileSize-h),
	)

	screen.DrawImage(frame, op)
}

func (c *CatSprite) getRow(dir domain.Direction) int {

	switch dir {
	case domain.DirDown:
		return 0
	case domain.DirUp:
		return 1
	case domain.DirLeft:
		return 2
	case domain.DirRight:
		return 3
	default:
		return 0
	}
}
