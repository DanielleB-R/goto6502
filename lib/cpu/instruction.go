package cpu

// Instruction data

type OperandFn func(*Processor) int
type OperationFn func(*Processor, int)

// Instruction contains the details for a single 6502 instruction
type Instruction struct {
	Opcode    byte
	Operand   OperandFn
	Operation OperationFn
	Length    int
}

// Execute executes the instruction against the current state of the CPU
func (i *Instruction) Execute(cpu *Processor) {
	i.Operation(cpu, i.Operand(cpu))
}

func asSigned(n byte) int8 {
	if n&0x80 == 0 {
		return int8(n)
	}
	return -int8(^n) - 1
}

// Address functions

func ImmediateAddress(p *Processor) int {
	return p.PC + 1
}

// ImmediateOperand gets the immediate byte operand of the current instruction
func ImmediateOperand(p *Processor) byte {
	return p.Memory.Read(ImmediateAddress(p))
}

func ZeroPageAddress(p *Processor) int {
	return int(ImmediateOperand(p))
}

func ZeroPageXAddress(p *Processor) int {
	return int(ImmediateOperand(p) + p.X)
}

func ZeroPageYAddress(p *Processor) int {
	return int(ImmediateOperand(p) + p.Y)
}

func AddressOperand(p *Processor) int {
	return p.Memory.ReadWord(ImmediateAddress(p))
}

func AbsoluteXAddress(p *Processor) int {
	return AddressOperand(p) + int(p.X)
}

func AbsoluteYAddress(p *Processor) int {
	return AddressOperand(p) + int(p.Y)
}

func IndirectAddress(p *Processor) int {
	return p.Memory.ReadWord(AddressOperand(p))
}

func IndexedIndirectAddress(p *Processor) int {
	addr := (ZeroPageAddress(p) + int(p.X)) & 0xff
	return p.Memory.ReadWord(addr)
}

func IndirectIndexedAddress(p *Processor) int {
	return p.Memory.ReadWord(ZeroPageAddress(p)) + int(p.Y)
}

func NoAddress(p *Processor) int {
	return 0
}

// Ops6502 is a 6502 opcode map
var Ops6502 = map[byte]Instruction{
	0x05: Instruction{0x05, ZeroPageAddress, ORA, 2},
	0x09: Instruction{0x09, ImmediateAddress, ORA, 2},
	0x0a: Instruction{0x0a, NoAddress, ASL, 1},
	0x0d: Instruction{0x0d, AddressOperand, ORA, 3},

	0x10: Instruction{0x10, ImmediateAddress, BPL, 2},
	0x15: Instruction{0x15, ZeroPageXAddress, ORA, 2},
	0x18: Instruction{0x18, NoAddress, CLC, 1},
	0x19: Instruction{0x19, AbsoluteYAddress, ORA, 3},
	0x1d: Instruction{0x1d, AbsoluteXAddress, ORA, 3},

	0x29: Instruction{0x29, ImmediateAddress, AND, 2},
	0x2a: Instruction{0x2a, NoAddress, ROL, 1},
	0x2d: Instruction{0x2d, AddressOperand, AND, 3},

	0x30: Instruction{0x30, ImmediateAddress, BMI, 2},
	0x38: Instruction{0x38, NoAddress, SEC, 1},
	0x39: Instruction{0x39, AbsoluteYAddress, AND, 3},
	0x3d: Instruction{0x3d, AbsoluteXAddress, AND, 3},

	0x4a: Instruction{0x4a, NoAddress, LSR, 1},
	0x4c: Instruction{0x4c, AddressOperand, JMP, 3},

	0x6a: Instruction{0x6a, NoAddress, ROR, 1},
	0x6c: Instruction{0x6c, IndirectAddress, JMP, 3},

	0x84: Instruction{0x84, ZeroPageAddress, STY, 2},
	0x85: Instruction{0x85, ZeroPageAddress, STA, 2},
	0x86: Instruction{0x86, ZeroPageAddress, STX, 2},
	0x88: Instruction{0x88, NoAddress, DEY, 1},
	0x8c: Instruction{0x8c, AddressOperand, STY, 3},
	0x8d: Instruction{0x8d, AddressOperand, STA, 3},
	0x8e: Instruction{0x8e, AddressOperand, STX, 3},

	0x90: Instruction{0x90, ImmediateAddress, BCC, 2},
	0x94: Instruction{0x94, ZeroPageXAddress, STY, 2},
	0x95: Instruction{0x95, ZeroPageXAddress, STA, 2},
	0x96: Instruction{0x96, ZeroPageYAddress, STX, 2},
	0x99: Instruction{0x99, AbsoluteYAddress, STA, 3},
	0x9d: Instruction{0x9d, AbsoluteXAddress, STA, 3},

	0xa0: Instruction{0xa0, ImmediateAddress, LDY, 2},
	0xa1: Instruction{0xa1, IndexedIndirectAddress, LDA, 2},
	0xa2: Instruction{0xa2, ImmediateAddress, LDX, 2},
	0xa4: Instruction{0xa4, ZeroPageAddress, LDY, 2},
	0xa5: Instruction{0xa5, ZeroPageAddress, LDA, 2},
	0xa6: Instruction{0xa6, ZeroPageAddress, LDX, 2},
	0xa8: Instruction{0xa8, NoAddress, TAY, 1},
	0xa9: Instruction{0xa9, ImmediateAddress, LDA, 2},
	0xaa: Instruction{0xaa, NoAddress, TAX, 1},
	0xac: Instruction{0xac, AddressOperand, LDY, 3},
	0xad: Instruction{0xad, AddressOperand, LDA, 3},
	0xae: Instruction{0xae, AddressOperand, LDX, 3},

	0xb0: Instruction{0xb0, ImmediateAddress, BCS, 2},
	0xb4: Instruction{0xb4, ZeroPageXAddress, LDY, 2},
	0xb5: Instruction{0xb5, ZeroPageXAddress, LDA, 2},
	0xb6: Instruction{0xb6, ZeroPageYAddress, LDX, 2},
	0xb9: Instruction{0xb9, AbsoluteYAddress, LDA, 3},
	0xbc: Instruction{0xbc, AbsoluteXAddress, LDY, 3},
	0xbd: Instruction{0xbd, AbsoluteXAddress, LDA, 3},
	0xbe: Instruction{0xbe, AbsoluteYAddress, LDX, 3},

	0xc6: Instruction{0xc6, ZeroPageAddress, DEC, 2},
	0xc8: Instruction{0xc8, NoAddress, INY, 1},
	0xca: Instruction{0xc8, NoAddress, DEX, 1},
	0xce: Instruction{0xce, AddressOperand, DEC, 3},

	0xd0: Instruction{0xd0, ImmediateAddress, BNE, 2},
	0xd6: Instruction{0xd6, ZeroPageXAddress, DEC, 2},
	0xde: Instruction{0xde, AbsoluteXAddress, DEC, 3},

	0xe6: Instruction{0xe6, ZeroPageAddress, INC, 2},
	0xe8: Instruction{0xe8, NoAddress, INX, 1},
	0xee: Instruction{0xee, AddressOperand, INC, 3},

	0xf0: Instruction{0xf0, ImmediateAddress, BEQ, 2},
	0xf6: Instruction{0xf6, ZeroPageXAddress, INC, 2},
	0xfe: Instruction{0xfe, AbsoluteXAddress, INC, 3},
}
