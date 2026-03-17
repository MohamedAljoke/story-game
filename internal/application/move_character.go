package application

import "story-game/internal/domain"

type MoveCharacter struct {
	World *domain.World
}

func (mc *MoveCharacter) Execute(characterID domain.CharacterID, dir domain.Direction) {
	mc.World.MoveCharacter(characterID, dir)
}
