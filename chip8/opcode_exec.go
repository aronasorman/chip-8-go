package chip8

func (c *C8) execOpcode(opcode uint16) *C8 {

	// the ANNN opcode instructs the CPU
	// to set the index to NNN
	if (opcode & 0xF000) == 0xA000 {
		c.setIndex(opcode & 0x0FFF)
	}

	// after executing the opcode, advance
	// the program counter by 2
	c.pc += 2

	return c
}

func (c *C8) setIndex(val uint16) {
	c.index = val
}
