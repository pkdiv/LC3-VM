package main

func Add(instr uint16) {
	dr := (instr >> 9) & 0x0007
	sr1 := (instr >> 6) & 0x0007
	mode := (instr >> 5) & 0x0001
	var operand2 uint16

	// IF mode == 0 register mode
	//         == 1 immediate mode
	if mode == 0 {
		sr2 := OP_CODE(instr & 7)
		operand2 = reg[sr2]
	} else {
		imd := instr & 31             // Immediate mode can have values from 0 to 2^5 - 1
		operand2 = signExtend(imd, 5) // Sign extend it
	}

	reg[dr] = reg[sr1] + operand2

	// Set the R_COND based on the value in reg[dr]
	updateFlag(dr)

}

func Ld(instr uint16) {
	dr := (instr >> 9) & 0x0007
	offset9 := instr & 0x01FF
	offset := signExtend(offset9, 9)
	reg[dr] = memory[reg[R_PC]+offset]
	updateFlag(dr)
}

func St(instr uint16) {
	sr := (instr >> 9) & 0x0007
	offset9 := instr & 0x01FF
	offset := signExtend(offset9, 9)

	memory[reg[R_PC]+offset] = reg[sr]
}

func And(instr uint16) {

	dr := (instr >> 9) & 0x0007
	sr1 := (instr >> 6) & 0x0007
	mode := (instr >> 5) & 0x0001
	var operand2 uint16

	if mode == 0 {
		sr2 := OP_CODE(instr & 7)
		operand2 = reg[sr2]
	} else {
		imd := instr & 0x001F         // Immediate mode can have values from 0 to 2^5 - 1
		operand2 = signExtend(imd, 5) // Sign extend it
	}

	reg[dr] = reg[sr1] & operand2

	updateFlag(dr)

}

func Ldr(instr uint16) {
	dr := (instr >> 9) & 0x0007
	base := (instr >> 6) & 0x0007

	offset6 := instr & 0x00EF
	offset := signExtend(offset6, 6)

	reg[dr] = memory[reg[base]+offset]

	updateFlag(dr)
}

func Str(instr uint16) {
	sr := (instr >> 9) & 0x0007
	base := (instr >> 6) & 0x0007

	offset6 := instr & 0x00EF
	offset := signExtend(offset6, 6)

	memory[reg[base]+offset] = reg[sr]

}

func Not(instr uint16) {
	dr := (instr >> 9) & 0x0007
	sr := (instr >> 6) & 0x0007

	reg[dr] = ^reg[sr]

	updateFlag(dr)

}

func Ldi(instr uint16) {
	dr := (instr >> 9) & 0x0007
	offset9 := instr & 0x01FF
	offset := signExtend(offset9, 9)
	reg[dr] = memory[memory[reg[R_PC]+offset]]
	updateFlag(dr)
}

func Sti(instr uint16) {
	sr := (instr >> 9) & 0x0007
	offset9 := instr & 0x01FF
	offset := signExtend(offset9, 9)

	memory[memory[reg[R_PC]+offset]] = reg[sr]

}

func updateFlag(register uint16) {
	if reg[register] == 0 {
		reg[R_COND] = uint16(FLAG_ZERO)
	} else if reg[register] > 0 {
		reg[R_COND] = uint16(FLAG_POS)
	} else {
		reg[R_COND] = uint16(FLAG_NEG)
	}
}

// Extend any bit value to 16 bits
func signExtend(value uint16, bitcount int) uint16 {

	significantBit := value >> uint16(bitcount-1)
	var extdValue uint16

	if significantBit == 0 {
		extdValue = 0x0000 | value
	} else {
		extdValue = (0xFFFF << bitcount) | value
	}

	return extdValue

}
