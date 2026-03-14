package engine

// Command represents a player input command.
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
