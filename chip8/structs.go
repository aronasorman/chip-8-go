package chip8

// C8 holds the state of a chip-8 CPU
type C8 struct {
	// CHIP-8 has 4K memory
	memory [4096]uint8

	// 16 registers
	registers [16]uint8

	// index register
	index uint16

	// program counter
	pc uint16

	// graphics memory
	gfx [64 * 32]uint8

	// timers
	delayTimer uint8
	soundTimer uint8

	// the stack and pointer
	stack uint16
	sp    uint16

	// the keypad
	keypad [16]uint8
}

// Initialize sets up the CHIP instance
func (c *C8) Initialize() {
}

// NewC8 creates a new CHIP-8 instance, with all
// systems initialized
func NewC8() *C8 {
	return &C8{}
}
