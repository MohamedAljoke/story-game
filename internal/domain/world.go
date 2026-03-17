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
	TileMap    *TileMap
}

func NewWorld(tm *TileMap) *World {
	w := &World{
		characters: make(map[CharacterID]*Character),
		positions:  make(map[CharacterID]*Position),
		TileMap:    tm,
	}

	spawnX, spawnY := w.findSpawn()

	w.characters[PlayerID] = &Character{ID: PlayerID, Name: "Player", Type: CharacterPlayer, Facing: DirDown}
	w.positions[PlayerID] = &Position{X: spawnX, Y: spawnY}

	w.characters[CatID] = &Character{ID: CatID, Name: "Cat", Type: CharacterPet, Leader: PlayerID}
	w.positions[CatID] = &Position{X: spawnX - TileSize, Y: spawnY}

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
	newX := pos.X + dx*MoveSpeed
	newY := pos.Y + dy*MoveSpeed

	if w.TileMap != nil {
		half := float64(TileSize) / 2
		cx := newX + half
		cy := newY + half
		if !w.TileMap.IsWalkable(cx, cy) {
			return
		}
	}

	pos.X = newX
	pos.Y = newY

	maxX := float64(WorldWidth - TileSize)
	maxY := float64(WorldHeight - TileSize)
	if w.TileMap != nil {
		maxX = float64(w.TileMap.PixelWidth() - TileSize)
		maxY = float64(w.TileMap.PixelHeight() - TileSize)
	}
	pos.Clamp(0, 0, maxX, maxY)
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

func (w *World) findSpawn() (float64, float64) {
	if w.TileMap != nil {
		sumX, sumY, count := 0, 0, 0
		for y := 0; y < w.TileMap.Height; y++ {
			for x := 0; x < w.TileMap.Width; x++ {
				tile := w.TileMap.TileAt(0, x, y)
				if tile != 0 {
					sumX += x
					sumY += y
					count++
				}
			}
		}
		if count > 0 {
			cx := float64(sumX/count) * float64(w.TileMap.TileSize)
			cy := float64(sumY/count) * float64(w.TileMap.TileSize)
			return cx, cy
		}
	}
	return float64(WorldWidth/2 - TileSize/2), float64(WorldHeight/2 - TileSize/2)
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
