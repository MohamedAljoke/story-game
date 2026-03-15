package application

import (
	"story-game/internal/domain"
	"story-game/internal/engine"
)

type GameLoop struct {
	MoveChar *MoveCharacter
	World    *domain.World
}

func (gl *GameLoop) ProcessCommands(cmds []engine.Command) bool {
	for _, cmd := range cmds {
		switch cmd {
		case engine.CmdQuit:
			return true
		case engine.CmdMoveUp:
			gl.MoveChar.Execute("player", domain.DirUp)
		case engine.CmdMoveDown:
			gl.MoveChar.Execute("player", domain.DirDown)
		case engine.CmdMoveLeft:
			gl.MoveChar.Execute("player", domain.DirLeft)
		case engine.CmdMoveRight:
			gl.MoveChar.Execute("player", domain.DirRight)
		}
	}

	// update followers every frame
	gl.World.UpdateFollowers()

	return false
}
