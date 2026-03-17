package ebiten

import (
	"image"
	"story-game/internal/adapter/ebiten/helpers"
	"story-game/internal/domain"
	"story-game/internal/loader"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type tilesetEntry struct {
	firstGID     int
	lastGID      int
	isCollection bool
	image        *eb.Image
	columns      int
	tileW        int
	tileH        int
	tiles        map[int]*eb.Image
	tileWidths   map[int]int
	tileHeights  map[int]int
}

type TileRenderer struct {
	entries []tilesetEntry
	scale   float64
}

func NewTileRenderer(infos []loader.TilesetInfo, scale float64) (*TileRenderer, error) {
	tr := &TileRenderer{scale: scale}

	for i, info := range infos {
		entry := tilesetEntry{
			firstGID:     info.FirstGID,
			isCollection: info.IsCollection,
		}

		nextFirstGID := 0
		if i+1 < len(infos) {
			nextFirstGID = infos[i+1].FirstGID
		}

		if info.IsCollection {
			entry.tiles = make(map[int]*eb.Image)
			entry.tileWidths = make(map[int]int)
			entry.tileHeights = make(map[int]int)

			for _, t := range info.Tiles {
				img, err := helpers.LoadSprite(t.Path)
				if err != nil {
					return nil, err
				}
				gid := info.FirstGID + t.LocalID
				entry.tiles[gid] = img
				entry.tileWidths[gid] = t.Width
				entry.tileHeights[gid] = t.Height
			}

			if nextFirstGID > 0 {
				entry.lastGID = nextFirstGID - 1
			} else {
				entry.lastGID = info.FirstGID + len(info.Tiles) - 1
			}
		} else {
			img, err := helpers.LoadSprite(info.ImagePath)
			if err != nil {
				return nil, err
			}
			entry.image = img
			entry.columns = info.Columns
			entry.tileW = info.TileW
			entry.tileH = info.TileH

			if nextFirstGID > 0 {
				entry.lastGID = nextFirstGID - 1
			} else {
				bounds := img.Bounds()
				rows := bounds.Dy() / info.TileH
				entry.lastGID = info.FirstGID + rows*info.Columns - 1
			}
		}

		tr.entries = append(tr.entries, entry)
	}

	return tr, nil
}

func (tr *TileRenderer) DrawMap(screen *eb.Image, tm *domain.TileMap, cam *domain.Camera) {
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

				tr.drawTile(screen, tm, cam, tileID, x, y)
			}
		}
	}
}

func (tr *TileRenderer) drawTile(screen *eb.Image, tm *domain.TileMap, cam *domain.Camera, tileID, x, y int) {
	for i := range tr.entries {
		e := &tr.entries[i]
		if tileID < e.firstGID || tileID > e.lastGID {
			continue
		}

		destX := float64(x*tm.TileSize) - cam.X
		destY := float64(y*tm.TileSize) - cam.Y

		if e.isCollection {
			img, ok := e.tiles[tileID]
			if !ok {
				return
			}
			op := &eb.DrawImageOptions{}
			op.GeoM.Scale(tr.scale, tr.scale)
			h := float64(e.tileHeights[tileID]) * tr.scale
			op.GeoM.Translate(destX, destY+float64(tm.TileSize)-h)
			screen.DrawImage(img, op)
		} else {
			idx := tileID - e.firstGID
			srcX := (idx % e.columns) * e.tileW
			srcY := (idx / e.columns) * e.tileH

			op := &eb.DrawImageOptions{}
			op.GeoM.Scale(tr.scale, tr.scale)
			op.GeoM.Translate(destX, destY)

			sub := e.image.SubImage(
				image.Rect(srcX, srcY, srcX+e.tileW, srcY+e.tileH),
			).(*eb.Image)
			screen.DrawImage(sub, op)
		}
		return
	}
}
