package chip8

import "testing"

func TestC8_fetchOpcode(t *testing.T) {
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

	var opCode1Mem [4096]uint8
	opCode1Mem[0] = 0xA2
	opCode1Mem[1] = 0xF0

	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		{
			name: "Check returned opcode",
			fields: fields{
				pc:     0,
				memory: opCode1Mem,
			},
			want: 0xA2F0,
		},
		// TODO: Add test cases.
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
			if got := c.fetchOpcode(); got != tt.want {
				t.Errorf("C8.fetchOpcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
