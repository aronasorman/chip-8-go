package chip8

import (
	"reflect"
	"testing"
)

func TestC8_execOpcode(t *testing.T) {
	type fields struct {
		DrawFlag   bool
		memory     [4096]uint8
		registers  [16]uint8
		index      uint16
		pc         uint16
		gfx        [64 * 32]uint8
		delayTimer uint8
		soundTimer uint8
		stack      *Stack
		sp         uint16
		keypad     [16]uint8
	}
	type args struct {
		opcode uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *C8
	}{
		// TODO: Add test cases.
		{
			name: "test ANNN opcode",
			args: args{
				opcode: 0xA2F0,
			},
			want: &C8{
				index: 0x02F0,
				pc:    2,
			},
		},
		{
			name: "test 1NNN opcode",
			args: args{
				opcode: 0x1FFF,
			},
			fields: fields{
				stack: NewStack(),
			},
			want: &C8{
				pc:    0xFFF,
				stack: NewStack(),
			},
		},
		{
			name: "test 00E0 opcode (clears the screen)",
			args: args{
				opcode: 0x00E0,
			},
			// doesn't do anything, aka screen should still be 0
			// and draw that
			want: &C8{
				pc:       2,
				DrawFlag: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &C8{
				DrawFlag:   tt.fields.DrawFlag,
				memory:     tt.fields.memory,
				registers:  tt.fields.registers,
				index:      tt.fields.index,
				pc:         tt.fields.pc,
				gfx:        tt.fields.gfx,
				delayTimer: tt.fields.delayTimer,
				soundTimer: tt.fields.soundTimer,
				stack:      tt.fields.stack,
				sp:         tt.fields.sp,
				keypad:     tt.fields.keypad,
			}
			if got := c.execOpcode(tt.args.opcode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("C8.execOpcode() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestC8_jumpTo(t *testing.T) {
	type fields struct {
		DrawFlag   bool
		memory     [4096]uint8
		registers  [16]uint8
		index      uint16
		pc         uint16
		gfx        [64 * 32]uint8
		delayTimer uint8
		soundTimer uint8
		stack      *Stack
		sp         uint16
		keypad     [16]uint8
	}
	type args struct {
		addr uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Jump sets pc to the set value",
			fields: fields{pc: 0x0, stack: NewStack()},
			args: args{
				addr: 0x900,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &C8{
				DrawFlag:   tt.fields.DrawFlag,
				memory:     tt.fields.memory,
				registers:  tt.fields.registers,
				index:      tt.fields.index,
				pc:         tt.fields.pc,
				gfx:        tt.fields.gfx,
				delayTimer: tt.fields.delayTimer,
				soundTimer: tt.fields.soundTimer,
				stack:      tt.fields.stack,
				sp:         tt.fields.sp,
				keypad:     tt.fields.keypad,
			}
			c.jumpTo(tt.args.addr)
			if pc := c.pc; pc != tt.args.addr {
				t.Errorf("pc was not set to the right value: %v", c.pc)
			}
		})
	}
}
