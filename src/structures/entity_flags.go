package structures

// Entity Flags

const (
	NoFlags         = 1 << iota // 1 << 0 = 1
	MovedFlag                   // 1 << 1 = 2
	StatChangedFlag             // 1 << 2 = 4
	// DeadFlag                    // 1 << 3 = 8 // mabye move Dead to a Flag
)
