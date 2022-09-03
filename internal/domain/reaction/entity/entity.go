package entity

// Command is entity for command.
type Command struct {
	Command  string
	ID       int
	Query    string
	Arg      string
	Page     int
	LastPage int
	Info     int8
}
