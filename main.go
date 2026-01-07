package main

// Max memory supported by LC3 VM 128KB
const MAX_MEMORY = 65536

type REG uint16
type OP_CODE uint16
type STATUS uint16

// Registers
const (
	R_R0 REG = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC   // Program Counter
	R_COND // Condition Register
	R_COUNT
)

const (
	OP_BR   OP_CODE = 0 //branch
	OP_ADD              // add
	OP_LD               // load
	OP_ST               // store
	OP_JSR              // jump register
	OP_AND              // bitwise and
	OP_LDR              // load register
	OP_STR              // store register
	OP_RTI              // unused
	OP_NOT              // bitwise not
	OP_LDI              // load indirect
	OP_STI              // store indirect
	OP_JMP              // jump
	OP_RES              // reserved (unused)
	OP_LEA              // load effective address
	OP_TRAP             // execute trap
)

const (
	FLAG_POS STATUS = 1 << iota
	FLAG_ZERO
	FLAG_NEG
)

var memory [MAX_MEMORY]uint16
var reg [R_COUNT]uint16

func main() {

	// Initialize Conditional Register to zero
	// PC to 0x3000
	reg[R_COND] = uint16(FLAG_ZERO)
	PC_START := 0x3000

	reg[R_PC] = uint16(PC_START)

	for {

		instr := memory[reg[R_PC]]
		reg[R_PC] = reg[R_PC] + 1

		// The 4 higher order bits store the OP_CODE of the instruction
		// The 12 remaining bits are interpreted according to the specific instruction.
		op_code := OP_CODE(instr >> 12)

		switch op_code {
		case OP_ADD:
			add(instr)
		}

	}

}

func add(instr uint16) {
	dr := (instr >> 9) & 7
	sr1 := (instr >> 6) & 7
	mode := (instr >> 5) & 1
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
	updateCondition(dr)

}

func updateCondition(register uint16) {
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
