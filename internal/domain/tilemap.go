package domain

type TileMap struct {
	Width    int
	Height   int
	TileSize int
	Layers   [][]int
}

func (tm *TileMap) TileAt(layer, x, y int) int {
	if layer < 0 || layer >= len(tm.Layers) {
		return 0
	}
	if x < 0 || x >= tm.Width || y < 0 || y >= tm.Height {
		return 0
	}
	return tm.Layers[layer][y*tm.Width+x]
}

func (tm *TileMap) PixelWidth() int {
	return tm.Width * tm.TileSize
}

func (tm *TileMap) PixelHeight() int {
	return tm.Height * tm.TileSize
}
