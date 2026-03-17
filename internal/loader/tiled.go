package loader

import (
	"encoding/json"
	"os"
	"story-game/internal/domain"
)

type tiledLayer struct {
	Data []int `json:"data"`
}

type tiledMap struct {
	Width      int          `json:"width"`
	Height     int          `json:"height"`
	TileWidth  int          `json:"tilewidth"`
	TileHeight int          `json:"tileheight"`
	Layers     []tiledLayer `json:"layers"`
}

func LoadTiledMap(path string) (*domain.TileMap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tm tiledMap
	if err := json.Unmarshal(data, &tm); err != nil {
		return nil, err
	}

	layers := make([][]int, len(tm.Layers))
	for i, l := range tm.Layers {
		layers[i] = l.Data
	}

	return &domain.TileMap{
		Width:    tm.Width,
		Height:   tm.Height,
		TileSize: tm.TileWidth,
		Layers:   layers,
	}, nil
}
