package domain

type CharacterType string

const (
	CharacterPlayer CharacterType = "player"
	CharacterPet    CharacterType = "pet"
	CharacterNPC    CharacterType = "npc"
)

type Character struct {
	ID     string
	Name   string
	Type   CharacterType
	Facing Direction
	Moving bool
}
