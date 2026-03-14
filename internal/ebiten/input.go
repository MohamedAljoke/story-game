package ebiten

import (
	"story-game/internal/engine"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type EbitenInput struct{}

func (e *EbitenInput) ReadCommands() []engine.Command {
	var cmds []engine.Command

	if eb.IsKeyPressed(eb.KeyW) || eb.IsKeyPressed(eb.KeyArrowUp) {
		cmds = append(cmds, engine.CmdMoveUp)
	}
	if eb.IsKeyPressed(eb.KeyS) || eb.IsKeyPressed(eb.KeyArrowDown) {
		cmds = append(cmds, engine.CmdMoveDown)
	}
	if eb.IsKeyPressed(eb.KeyA) || eb.IsKeyPressed(eb.KeyArrowLeft) {
		cmds = append(cmds, engine.CmdMoveLeft)
	}
	if eb.IsKeyPressed(eb.KeyD) || eb.IsKeyPressed(eb.KeyArrowRight) {
		cmds = append(cmds, engine.CmdMoveRight)
	}
	if inpututil.IsKeyJustPressed(eb.KeyEscape) {
		cmds = append(cmds, engine.CmdQuit)
	}

	return cmds
}
