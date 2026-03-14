package domain

const (
	WorldWidth  = 800
	WorldHeight = 600
	TileSize    = 32
	MoveSpeed   = 3.0
)

type World struct {
	Characters map[string]*Character
	Positions  map[string]*Position
}

func NewWorld() *World {
	w := &World{
		Characters: make(map[string]*Character),
		Positions:  make(map[string]*Position),
	}

	w.Characters["player"] = &Character{ID: "player", Name: "Player"}
	w.Positions["player"] = &Position{
		X: float64(WorldWidth/2 - TileSize/2),
		Y: float64(WorldHeight/2 - TileSize/2),
	}

	return w
}

func (w *World) MoveCharacter(id string, dir Direction) {
	pos, ok := w.Positions[id]
	if !ok {
		return
	}

	dx, dy := dir.Delta()
	pos.X += dx * MoveSpeed
	pos.Y += dy * MoveSpeed
	pos.Clamp(0, 0, WorldWidth-TileSize, WorldHeight-TileSize)
}
