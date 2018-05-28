package chip8

func (c *C8) execOpcode(opcode uint16) *C8 {

	//
	// next opcodes don't have any arguments
	// compare straight against their full value
	switch opcode {
	case 0x00E0:
		c.clearScreen()
		c.DrawFlag = true
	case 0x00EE:
		c.popStack()
	}

	// the other opcodes have arguments in them,
	// so get the first byte to determine
	// what opcode it is
	op := opcode >> 12

	switch op {
	// the ANNN opcode instructs the CPU
	// to set the index to NNN
	case 0xA:
		c.setIndex(opcode & 0x0FFF)
	// the 1NNN opcode jumps to an address,
	// without adding to the stack
	case 0x1:
		c.jumpTo(opcode & 0x0FFF)
		// subtract two, to take into account the
		// other +2 increment at the end of this function
		c.pc -= 2
	}

	// after executing the opcode, advance
	// the program counter by 2
	c.pc += 2

	return c
}

func (c *C8) setIndex(val uint16) {
	c.index = val
}

func (c *C8) clearScreen() {
	var newGfx [64 * 32]uint8
	c.gfx = newGfx
}

// jumpTo records the current program counter to
// the stack, then sets the program counter to the
// given address
func (c *C8) jumpTo(addr uint16) {
	c.pc = addr
}

// popStack sets the program counter to the last program
// counter valuye before a jump
func (c *C8) popStack() {
	pc, err := c.stack.Pop()
	if err != nil {
		panic("empty stack popped")
	}

	c.pc = pc
}
