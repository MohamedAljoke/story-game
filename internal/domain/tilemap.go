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

func (tm *TileMap) IsWalkable(px, py float64) bool {
	tx := int(px) / tm.TileSize
	ty := int(py) / tm.TileSize
	return tm.TileAt(0, tx, ty) != 0
}

func (tm *TileMap) PixelWidth() int {
	return tm.Width * tm.TileSize
}

func (tm *TileMap) PixelHeight() int {
	return tm.Height * tm.TileSize
}
