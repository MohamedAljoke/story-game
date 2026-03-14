package engine

type Command int

const (
	CmdMoveUp Command = iota
	CmdMoveDown
	CmdMoveLeft
	CmdMoveRight
	CmdQuit
)

type InputReader interface {
	ReadCommands() []Command
}
