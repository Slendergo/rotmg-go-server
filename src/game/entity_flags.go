package game

// Entity Flags
const (
	None        = 1 << iota // 1 << 0 = 1
	Moved                   // 1 << 1 = 2
	StatChanged             // 1 << 2 = 4
)
