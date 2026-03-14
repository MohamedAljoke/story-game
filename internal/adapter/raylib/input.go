package raylib

import (
	"bufio"
	"os"
	"story-game/internal/engine"
)

type RaylibInput struct {
	scanner *bufio.Scanner
}

func NewRaylibInput() *RaylibInput {
	return &RaylibInput{scanner: bufio.NewScanner(os.Stdin)}
}

func (r *RaylibInput) ReadCommands() []engine.Command {
	if !r.scanner.Scan() {
		return nil
	}

	var cmds []engine.Command
	for _, ch := range r.scanner.Text() {
		switch ch {
		case 'w', 'W':
			cmds = append(cmds, engine.CmdMoveUp)
		case 's', 'S':
			cmds = append(cmds, engine.CmdMoveDown)
		case 'a', 'A':
			cmds = append(cmds, engine.CmdMoveLeft)
		case 'd', 'D':
			cmds = append(cmds, engine.CmdMoveRight)
		case 'q', 'Q':
			cmds = append(cmds, engine.CmdQuit)
		}
	}
	return cmds
}
