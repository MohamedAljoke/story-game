package loader

import (
	"encoding/json"
	"os"
	"path/filepath"
	"story-game/internal/domain"
)

type tiledLayer struct {
	Data []int `json:"data"`
}

type tiledTilesetRef struct {
	FirstGID int    `json:"firstgid"`
	Source   string `json:"source"`
}

type tiledMap struct {
	Width      int               `json:"width"`
	Height     int               `json:"height"`
	TileWidth  int               `json:"tilewidth"`
	TileHeight int               `json:"tileheight"`
	Layers     []tiledLayer      `json:"layers"`
	Tilesets   []tiledTilesetRef `json:"tilesets"`
}

type tiledTileset struct {
	Columns    int               `json:"columns"`
	Image      string            `json:"image"`
	TileWidth  int               `json:"tilewidth"`
	TileHeight int               `json:"tileheight"`
	TileCount  int               `json:"tilecount"`
	Tiles      []tiledTileEntry  `json:"tiles"`
}

type tiledTileEntry struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	ImageWidth  int    `json:"imagewidth"`
	ImageHeight int    `json:"imageheight"`
}

type TilesetInfo struct {
	FirstGID     int
	IsCollection bool
	ImagePath    string
	TileW        int
	TileH        int
	Columns      int
	Tiles        []CollectionTile
}

type CollectionTile struct {
	LocalID int
	Path    string
	Width   int
	Height  int
}

func LoadTiledMap(path string, scale int) (*domain.TileMap, []TilesetInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	var tm tiledMap
	if err := json.Unmarshal(data, &tm); err != nil {
		return nil, nil, err
	}

	mapDir := filepath.Dir(path)

	layers := make([][]int, len(tm.Layers))
	for i, l := range tm.Layers {
		layers[i] = l.Data
	}

	var tilesets []TilesetInfo
	for _, ref := range tm.Tilesets {
		tsPath := filepath.Join(mapDir, ref.Source)
		tsData, err := os.ReadFile(tsPath)
		if err != nil {
			return nil, nil, err
		}

		var ts tiledTileset
		if err := json.Unmarshal(tsData, &ts); err != nil {
			return nil, nil, err
		}

		tsDir := filepath.Dir(tsPath)

		if ts.Columns > 0 {
			tilesets = append(tilesets, TilesetInfo{
				FirstGID:  ref.FirstGID,
				ImagePath: filepath.Clean(filepath.Join(tsDir, ts.Image)),
				TileW:     ts.TileWidth,
				TileH:     ts.TileHeight,
				Columns:   ts.Columns,
			})
		} else {
			var tiles []CollectionTile
			for _, t := range ts.Tiles {
				tiles = append(tiles, CollectionTile{
					LocalID: t.ID,
					Path:    filepath.Clean(filepath.Join(tsDir, t.Image)),
					Width:   t.ImageWidth,
					Height:  t.ImageHeight,
				})
			}
			tilesets = append(tilesets, TilesetInfo{
				FirstGID:     ref.FirstGID,
				IsCollection: true,
				Tiles:        tiles,
			})
		}
	}

	return &domain.TileMap{
		Width:    tm.Width,
		Height:   tm.Height,
		TileSize: tm.TileWidth * scale,
		Layers:   layers,
	}, tilesets, nil
}
