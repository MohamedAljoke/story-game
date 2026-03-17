package ebiten

import (
	"image"
	"story-game/internal/adapter/ebiten/helpers"
	"story-game/internal/domain"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Tileset struct {
	image   *eb.Image
	columns int
	tileW   int
	tileH   int
}

func NewTileset(path string, tileW, tileH int) (*Tileset, error) {
	img, err := helpers.LoadSprite(path)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	return &Tileset{
		image:   img,
		columns: bounds.Dx() / tileW,
		tileW:   tileW,
		tileH:   tileH,
	}, nil
}

func (ts *Tileset) DrawMap(screen *eb.Image, tm *domain.TileMap) {
	for layerIdx := range tm.Layers {
		for y := 0; y < tm.Height; y++ {
			for x := 0; x < tm.Width; x++ {
				tileID := tm.TileAt(layerIdx, x, y)
				if tileID == 0 {
					continue
				}

				idx := tileID - 1
				srcX := (idx % ts.columns) * ts.tileW
				srcY := (idx / ts.columns) * ts.tileH

				op := &eb.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tm.TileSize), float64(y*tm.TileSize))

				sub := ts.image.SubImage(
					image.Rect(srcX, srcY, srcX+ts.tileW, srcY+ts.tileH),
				).(*eb.Image)

				screen.DrawImage(sub, op)
			}
		}
	}
}
