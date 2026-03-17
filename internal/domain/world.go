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

	w.Characters[CatID] = &Character{ID: CatID, Name: "Cat", Type: CharacterPet, Leader: PlayerID}
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
	for _, char := range w.Characters {
		if char.Leader == "" {
			continue
		}

		leaderPos := w.Positions[char.Leader]
		followerPos := w.Positions[char.ID]

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
func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
