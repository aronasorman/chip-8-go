package main

import (
	"fmt"

	"github.com/aronasorman/chip-8-go/chip8"
)

// CHIP-8 has 4K memory
var memory [4096]uint8

// 16 registers
var registers [16]uint8

// index register
var index uint16

// program counter
var pc uint16

// graphics memory
var gfx [64 * 32]uint8

// timers
var delayTimer uint8
var soundTimer uint8

// the stack and pointer
var stack uint16
var sp uint16

// the keypad
var keypad [16]uint8

func main() {
	fmt.Println("Hello CHIP-8!")
	chip := chip8.NewC8()
}
