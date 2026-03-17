package domain

const (
	WorldWidth  = 800
	WorldHeight = 600
	TileSize    = 32
	MoveSpeed   = 3.0
)

type World struct {
	Characters map[CharacterID]*Character
	Positions  map[CharacterID]*Position
}

func NewWorld() *World {
	w := &World{
		Characters: make(map[CharacterID]*Character),
		Positions:  make(map[CharacterID]*Position),
	}

	w.Characters[PlayerID] = &Character{ID: PlayerID, Name: "Player", Type: CharacterPlayer, Facing: DirDown}
	w.Positions[PlayerID] = &Position{
		X: float64(WorldWidth/2 - TileSize/2),
		Y: float64(WorldHeight/2 - TileSize/2),
	}

	w.Characters[CatID] = &Character{ID: CatID, Name: "Cat", Type: CharacterPet}
	w.Positions[CatID] = &Position{
		X: float64(WorldWidth/2 - TileSize/2 - TileSize),
		Y: float64(WorldHeight/2 - TileSize/2),
	}

	return w
}

func (w *World) MoveCharacter(id CharacterID, dir Direction) {
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

	playerPos := w.Positions[PlayerID]
	catPos := w.Positions[CatID]

	if playerPos == nil || catPos == nil {
		return
	}

	cat := w.Characters[CatID]

	dx := playerPos.X - catPos.X
	dy := playerPos.Y - catPos.Y

	if abs(dx) < TileSize && abs(dy) < TileSize {
		cat.Moving = false
		return
	}

	cat.Moving = true

	// try to align with player plane first
	if abs(dx) > 2 {

		if dx > 0 {
			catPos.X += MoveSpeed
			cat.Facing = DirRight
		} else {
			catPos.X -= MoveSpeed
			cat.Facing = DirLeft
		}

		return
	}

	// once aligned horizontally, move vertically
	if abs(dy) > 2 {

		if dy > 0 {
			catPos.Y += MoveSpeed
			cat.Facing = DirDown
		} else {
			catPos.Y -= MoveSpeed
			cat.Facing = DirUp
		}
	}
}
func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
