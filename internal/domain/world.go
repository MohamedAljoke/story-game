package domain

const (
	WorldWidth  = 800
	WorldHeight = 600
	TileSize    = 32
	MoveSpeed   = 3.0
)

type World struct {
	characters map[CharacterID]*Character
	positions  map[CharacterID]*Position
}

func NewWorld() *World {
	w := &World{
		characters: make(map[CharacterID]*Character),
		positions:  make(map[CharacterID]*Position),
	}

	w.characters[PlayerID] = &Character{ID: PlayerID, Name: "Player", Type: CharacterPlayer, Facing: DirDown}
	w.positions[PlayerID] = &Position{
		X: float64(WorldWidth/2 - TileSize/2),
		Y: float64(WorldHeight/2 - TileSize/2),
	}

	w.characters[CatID] = &Character{ID: CatID, Name: "Cat", Type: CharacterPet, Leader: PlayerID}
	w.positions[CatID] = &Position{
		X: float64(WorldWidth/2 - TileSize/2 - TileSize),
		Y: float64(WorldHeight/2 - TileSize/2),
	}

	return w
}

func (w *World) MoveCharacter(id CharacterID, dir Direction) {
	pos, ok := w.positions[id]
	if !ok {
		return
	}

	char, ok := w.characters[id]
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
	for _, char := range w.characters {
		if char.Leader == "" {
			continue
		}

		leaderPos := w.positions[char.Leader]
		followerPos := w.positions[char.ID]

		if leaderPos == nil || followerPos == nil {
			continue
		}

		dx := leaderPos.X - followerPos.X
		dy := leaderPos.Y - followerPos.Y

		if abs(dx) < TileSize && abs(dy) < TileSize {
			char.Moving = false
			continue
		}

		char.Moving = true

		if abs(dx) > 2 {
			if dx > 0 {
				followerPos.X += MoveSpeed
				char.Facing = DirRight
			} else {
				followerPos.X -= MoveSpeed
				char.Facing = DirLeft
			}
			continue
		}

		if abs(dy) > 2 {
			if dy > 0 {
				followerPos.Y += MoveSpeed
				char.Facing = DirDown
			} else {
				followerPos.Y -= MoveSpeed
				char.Facing = DirUp
			}
		}
	}
}

func (w *World) Character(id CharacterID) *Character { return w.characters[id] }
func (w *World) Position(id CharacterID) *Position   { return w.positions[id] }

func (w *World) EachCharacter(fn func(CharacterID, *Character, *Position)) {
	for id, char := range w.characters {
		fn(id, char, w.positions[id])
	}
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
