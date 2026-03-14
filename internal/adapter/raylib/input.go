package raylib

import (
	"story-game/internal/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RaylibInput struct{}

func (r *RaylibInput) ReadCommands() []engine.Command {
	var cmds []engine.Command

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		cmds = append(cmds, engine.CmdMoveUp)
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		cmds = append(cmds, engine.CmdMoveDown)
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		cmds = append(cmds, engine.CmdMoveLeft)
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		cmds = append(cmds, engine.CmdMoveRight)
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		cmds = append(cmds, engine.CmdQuit)
	}

	return cmds
}
