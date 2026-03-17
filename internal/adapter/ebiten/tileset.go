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
	scale   float64
}

func NewTileset(path string, tileW, tileH int, scale float64) (*Tileset, error) {
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
		scale:   scale,
	}, nil
}

func (ts *Tileset) DrawMap(screen *eb.Image, tm *domain.TileMap, cam *domain.Camera) {
	startX := int(cam.X) / tm.TileSize
	startY := int(cam.Y) / tm.TileSize
	endX := startX + int(cam.Width)/tm.TileSize + 2
	endY := startY + int(cam.Height)/tm.TileSize + 2

	if startX < 0 {
		startX = 0
	}
	if startY < 0 {
		startY = 0
	}
	if endX > tm.Width {
		endX = tm.Width
	}
	if endY > tm.Height {
		endY = tm.Height
	}

	for layerIdx := range tm.Layers {
		for y := startY; y < endY; y++ {
			for x := startX; x < endX; x++ {
				tileID := tm.TileAt(layerIdx, x, y)
				if tileID == 0 {
					continue
				}

				idx := tileID - 1
				srcX := (idx % ts.columns) * ts.tileW
				srcY := (idx / ts.columns) * ts.tileH

				op := &eb.DrawImageOptions{}
				op.GeoM.Scale(ts.scale, ts.scale)
				op.GeoM.Translate(float64(x*tm.TileSize)-cam.X, float64(y*tm.TileSize)-cam.Y)

				sub := ts.image.SubImage(
					image.Rect(srcX, srcY, srcX+ts.tileW, srcY+ts.tileH),
				).(*eb.Image)

				screen.DrawImage(sub, op)
			}
		}
	}
}
