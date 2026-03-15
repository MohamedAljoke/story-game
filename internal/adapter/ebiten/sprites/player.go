package sprites

import (
	"story-game/internal/adapter/ebiten/helpers"
	"story-game/internal/domain"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type PlayerSprite struct {
	frames      [][]*eb.Image
	animFrame   int
	animCounter int
}

func NewPlayerSprite() *PlayerSprite {
	sheet, err := helpers.LoadSprite("assets/sprites/player.png")
	if err != nil {
		panic(err)
	}

	frames := helpers.SliceSpriteSheet(sheet, 4, 4, 48, 48)

	return &PlayerSprite{
		frames: frames,
	}
}

func (p *PlayerSprite) Draw(screen *eb.Image, pos *domain.Position, dir domain.Direction) {

	p.animCounter++

	if p.animCounter > 60 {
		p.animFrame = (p.animFrame + 1) % 4
		p.animCounter = 0
	}

	frame := p.getFrameFromRow(dir)

	op := &eb.DrawImageOptions{}
	op.GeoM.Translate(pos.X-8, pos.Y-16)

	screen.DrawImage(frame, op)
}

func (p *PlayerSprite) getFrameFromRow(dir domain.Direction) *eb.Image {

	row := 0

	switch dir {
	case domain.DirDown:
		row = 0
	case domain.DirUp:
		row = 1
	case domain.DirLeft:
		row = 2
	case domain.DirRight:
		row = 3
	}

	return p.frames[row][p.animFrame]
}
