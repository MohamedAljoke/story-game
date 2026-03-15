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

	w.Characters["player"] = &Character{ID: "player", Name: "Player", Type: CharacterPlayer, Facing: DirDown}
	w.Positions["player"] = &Position{
		X: float64(WorldWidth/2 - TileSize/2),
		Y: float64(WorldHeight/2 - TileSize/2),
	}

	w.Characters["cat"] = &Character{ID: "cat", Name: "Cat", Type: CharacterPet}
	w.Positions["cat"] = &Position{
		X: float64(WorldWidth/2 - TileSize/2 - TileSize),
		Y: float64(WorldHeight/2 - TileSize/2),
	}

	return w
}

func (w *World) MoveCharacter(id string, dir Direction) {
	pos, ok := w.Positions[id]
	if !ok {
		return
	}

	char, ok := w.Characters[id]
	if ok {
		char.Facing = dir
		char.Moving = true
	}

	dx, dy := dir.Delta()
	pos.X += dx * MoveSpeed
	pos.Y += dy * MoveSpeed
	pos.Clamp(0, 0, WorldWidth-TileSize, WorldHeight-TileSize)
}

func (w *World) UpdateFollowers() {

	playerPos := w.Positions["player"]
	catPos := w.Positions["cat"]

	if playerPos == nil || catPos == nil {
		return
	}

	dx := playerPos.X - catPos.X
	dy := playerPos.Y - catPos.Y

	// distance check so cat stays one tile behind
	if abs(dx) < TileSize && abs(dy) < TileSize {
		return
	}

	if dx > 0 {
		catPos.X += MoveSpeed
	}
	if dx < 0 {
		catPos.X -= MoveSpeed
	}

	if dy > 0 {
		catPos.Y += MoveSpeed
	}
	if dy < 0 {
		catPos.Y -= MoveSpeed
	}
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
