package port_scanner

type State int

const (
	Unknown = iota
	Open    = 1
	Closed  = 2
)
