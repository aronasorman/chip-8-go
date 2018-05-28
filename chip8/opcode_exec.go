package chip8

func (c *C8) execOpcode(opcode uint16) *C8 {

	//
	// next opcodes don't have any arguments
	// compare straight against their full value
	switch opcode {
	case 0x00E0:
		c.clearScreen()
		c.DrawFlag = true
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
	case 0x1:
		c.jumpTo(opcode & 0x0FFF)
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
	c.gfx = newGfx
}

// jumpTo records the current program counter to
// the stack, then sets the program counter to the
// given address
func (c *C8) jumpTo(addr uint16) {
	c.stack.Push(c.pc)
	c.pc = addr
}

