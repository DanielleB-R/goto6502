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
	return int(ImmediateOperand(p)+p.X) & 0xff
}

func ZeroPageYAddress(p *Processor) int {
	return int(ImmediateOperand(p)+p.Y) & 0xff
}

func AddressOperand(p *Processor) int {
	return p.Memory.ReadWord(ImmediateAddress(p))
}

func AbsoluteXAddress(p *Processor) int {
	return (AddressOperand(p) + int(p.X)) & 0xffff
}

func AbsoluteYAddress(p *Processor) int {
	return (AddressOperand(p) + int(p.Y)) & 0xffff
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

func ImmediateInstruction(opcode byte, operation OperationFn) Instruction {
	return Instruction{
		Opcode:    opcode,
		Operand:   ImmediateAddress,
		Operation: operation,
		Length:    2,
	}
}

func ZeroPageInstruction(opcode byte, operation OperationFn) Instruction {
	return Instruction{
		Opcode:    opcode,
		Operand:   ZeroPageAddress,
		Operation: operation,
		Length:    2,
	}
}

func ZeroPageXInstruction(opcode byte, operation OperationFn) Instruction {
	return Instruction{
		Opcode:    opcode,
		Operand:   ZeroPageXAddress,
		Operation: operation,
		Length:    2,
	}
}

func ZeroPageYInstruction(opcode byte, operation OperationFn) Instruction {
	return Instruction{
		Opcode:    opcode,
		Operand:   ZeroPageYAddress,
		Operation: operation,
		Length:    2,
	}
}

func AbsoluteInstruction(opcode byte, operation OperationFn) Instruction {
	return Instruction{
		Opcode:    opcode,
		Operand:   AddressOperand,
		Operation: operation,
		Length:    3,
	}
}

func AbsoluteXInstruction(opcode byte, operation OperationFn) Instruction {
	return Instruction{
		Opcode:    opcode,
		Operand:   AbsoluteXAddress,
		Operation: operation,
		Length:    3,
	}
}

func AbsoluteYInstruction(opcode byte, operation OperationFn) Instruction {
	return Instruction{
		Opcode:    opcode,
		Operand:   AbsoluteYAddress,
		Operation: operation,
		Length:    3,
	}
}

func NoOperandInstruction(opcode byte, operation OperationFn) Instruction {
	return Instruction{
		Opcode:    opcode,
		Operand:   NoAddress,
		Operation: operation,
		Length:    1,
	}
}

// Ops6502 is a 6502 opcode map
var Ops6502 = map[byte]Instruction{
	0x05: ZeroPageInstruction(0x05, ORA),
	0x09: ImmediateInstruction(0x09, ORA),
	0x0a: NoOperandInstruction(0x0a, ASL),
	0x0d: AbsoluteInstruction(0x0d, ORA),

	0x10: ImmediateInstruction(0x10, BPL),
	0x15: ZeroPageXInstruction(0x15, ORA),
	0x18: NoOperandInstruction(0x18, CLC),
	0x19: AbsoluteYInstruction(0x19, ORA),
	0x1d: AbsoluteXInstruction(0x1d, ORA),

	0x21: Instruction{0x21, IndexedIndirectAddress, AND, 2},
	0x29: ImmediateInstruction(0x29, AND),
	0x2a: NoOperandInstruction(0x2a, ROL),
	0x2d: AbsoluteInstruction(0x2d, AND),

	0x30: ImmediateInstruction(0x30, BMI),
	0x31: Instruction{0x31, IndirectIndexedAddress, AND, 2},
	0x38: NoOperandInstruction(0x38, SEC),
	0x39: AbsoluteYInstruction(0x39, AND),
	0x3d: AbsoluteXInstruction(0x3d, AND),

	0x48: NoOperandInstruction(0x48, PHA),
	0x4a: NoOperandInstruction(0x4a, LSR),
	0x4c: AbsoluteInstruction(0x4c, JMP),

	0x68: NoOperandInstruction(0x68, PLA),
	0x6a: NoOperandInstruction(0x6a, ROR),
	0x6c: Instruction{0x6c, IndirectAddress, JMP, 3},

	0x81: Instruction{0x81, IndexedIndirectAddress, STA, 2},
	0x84: ZeroPageInstruction(0x84, STY),
	0x85: ZeroPageInstruction(0x85, STA),
	0x86: ZeroPageInstruction(0x86, STX),
	0x88: NoOperandInstruction(0x88, DEY),
	0x8a: NoOperandInstruction(0x8a, TXA),
	0x8c: AbsoluteInstruction(0x8c, STY),
	0x8d: AbsoluteInstruction(0x8d, STA),
	0x8e: AbsoluteInstruction(0x8e, STX),

	0x90: ImmediateInstruction(0x90, BCC),
	0x91: Instruction{0x91, IndirectIndexedAddress, STA, 2},
	0x94: ZeroPageXInstruction(0x94, STY),
	0x95: ZeroPageXInstruction(0x95, STA),
	0x96: ZeroPageYInstruction(0x96, STX),
	0x98: NoOperandInstruction(0x98, TYA),
	0x99: AbsoluteYInstruction(0x99, STA),
	0x9a: NoOperandInstruction(0x9a, TXS),
	0x9d: AbsoluteXInstruction(0x9d, STA),

	0xa0: ImmediateInstruction(0xa0, LDY),
	0xa1: Instruction{0xa1, IndexedIndirectAddress, LDA, 2},
	0xa2: ImmediateInstruction(0xa2, LDX),
	0xa4: ZeroPageInstruction(0xa4, LDY),
	0xa5: ZeroPageInstruction(0xa5, LDA),
	0xa6: ZeroPageInstruction(0xa6, LDX),
	0xa8: NoOperandInstruction(0xa8, TAY),
	0xa9: ImmediateInstruction(0xa9, LDA),
	0xaa: NoOperandInstruction(0xaa, TAX),
	0xac: AbsoluteInstruction(0xac, LDY),
	0xad: AbsoluteInstruction(0xad, LDA),
	0xae: AbsoluteInstruction(0xae, LDX),

	0xb0: ImmediateInstruction(0xb0, BCS),
	0xb1: Instruction{0xb1, IndirectIndexedAddress, LDA, 2},
	0xb4: ZeroPageXInstruction(0xb4, LDY),
	0xb5: ZeroPageXInstruction(0xb5, LDA),
	0xb6: ZeroPageYInstruction(0xb6, LDX),
	0xb9: AbsoluteYInstruction(0xb9, LDA),
	0xba: NoOperandInstruction(0xba, TSX),
	0xbc: AbsoluteXInstruction(0xbc, LDY),
	0xbd: AbsoluteXInstruction(0xbd, LDA),
	0xbe: AbsoluteYInstruction(0xbe, LDX),

	0xc6: ZeroPageInstruction(0xc6, DEC),
	0xc8: NoOperandInstruction(0xc8, INY),
	0xca: NoOperandInstruction(0xc8, DEX),
	0xce: AbsoluteInstruction(0xce, DEC),

	0xd0: ImmediateInstruction(0xd0, BNE),
	0xd6: ZeroPageXInstruction(0xd6, DEC),
	0xde: AbsoluteXInstruction(0xde, DEC),

	0xe6: ZeroPageInstruction(0xe6, INC),
	0xe8: NoOperandInstruction(0xe8, INX),
	0xea: NoOperandInstruction(0xea, NOP),
	0xee: AbsoluteInstruction(0xee, INC),

	0xf0: ImmediateInstruction(0xf0, BEQ),
	0xf6: ZeroPageXInstruction(0xf6, INC),
	0xfe: AbsoluteXInstruction(0xfe, INC),
}
