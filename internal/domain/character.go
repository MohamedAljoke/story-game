package domain

type CharacterID string

const (
	PlayerID CharacterID = "player"
	CatID    CharacterID = "cat"
)

type CharacterType string

const (
	CharacterPlayer CharacterType = "player"
	CharacterPet    CharacterType = "pet"
	CharacterNPC    CharacterType = "npc"
)

type Character struct {
	ID     CharacterID
	Name   string
	Type   CharacterType
	Facing Direction
	Moving bool
	Leader CharacterID
}
