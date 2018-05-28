package chip8

// C8 holds the state of a chip-8 CPU
type C8 struct {
	// DrawFlag if set, tells the invoker
	// that something is ready to draw
	DrawFlag bool

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
	stack *Stack
	sp    uint16

	// the keypad
	keypad [16]uint8
}

// LoadGame loads the ROM on the given
// path
func (c *C8) LoadGame(path string) {

}

// GetKeyInput reads the input from the
// keyboard or buttons, and sets them as input
// for the next cycle
func (c *C8) GetKeyInput() {

}

// EmulateCycle runs the chip for one cycle, taking
// in any opcodes defined in the game, and executes
// them
func (c *C8) EmulateCycle() {
	opcode := c.fetchOpcode()
	c.executeOpcode(opcode)
}

// Draw draws the current framebuffer
func (c *C8) Draw() {

}

// NewC8 creates a new CHIP-8 instance, with all
// systems initialized
func NewC8() *C8 {
	return &C8{
		pc:    0x200,
		stack: NewStack(),
	}

	// load fontset on to 0x80 in memory
}

// fetchOpcode gets the next opcode in memory, based on the
// current pc
func (c *C8) fetchOpcode() uint16 {
	// an opcode on the CHIP-8 is composed of two bytes
	// in continguous memory
	opcode1 := uint16(c.memory[c.pc])
	opcode2 := uint16(c.memory[c.pc+1])

	// merge the two bytes together
	fullOpcode := opcode1<<8 | opcode2
	return fullOpcode
}

func (c *C8) executeOpcode(opcode uint16) {
	c.execOpcode(opcode)
}
