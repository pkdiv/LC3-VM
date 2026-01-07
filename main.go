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
	OP_BR   OP_CODE = iota //branch
	OP_ADD                 // add
	OP_LD                  // load
	OP_ST                  // store
	OP_JSR                 // jump register
	OP_AND                 // bitwise and
	OP_LDR                 // load register
	OP_STR                 // store register
	OP_RTI                 // unused
	OP_NOT                 // bitwise not
	OP_LDI                 // load indirect
	OP_STI                 // store indirect
	OP_JMP                 // jump
	OP_RES                 // reserved (unused)
	OP_LEA                 // load effective address
	OP_TRAP                // execute trap
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
		reg[R_PC]++

		// The 4 higher order bits store the OP_CODE of the instruction
		// The 12 remaining bits are interpreted according to the specific instruction.
		op_code := OP_CODE(instr >> 12)

		switch op_code {
		case OP_BR:
			panic("Not Implemented")
		case OP_ADD:
			Add(instr)
		case OP_LD:
			Ld(instr)
		case OP_ST:
			St(instr)
		case OP_JSR:
			panic("Not Implemented")
		case OP_AND:
			And(instr)
		case OP_LDR:
			Ldr(instr)
		case OP_STR:
			Str(instr)
		case OP_RTI:
			panic("Not Implemented")
		case OP_NOT:
			Not(instr)
		case OP_LDI:
			Ldi(instr)
		case OP_STI:
			Sti(instr)
		case OP_JMP:
			panic("Not Implemented")
		case OP_RES:
			panic("Not Implemented")
		case OP_LEA:
			panic("Not Implemented")
		case OP_TRAP:
			panic("Not Implemented")
		}

	}

}
