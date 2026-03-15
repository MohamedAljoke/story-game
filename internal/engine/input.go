package engine

type Command int

const (
	CmdMoveUp Command = iota
	CmdMoveDown
	CmdMoveLeft
	CmdMoveRight
	CmdQuit
)
