package cpu

// Instruction data

type operandFn func(*Processor) int
type operationFn func(*Processor, int)

// Instruction contains the details for a single 6502 instruction
type Instruction struct {
	Operand   operandFn
	Operation operationFn
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

func ImmediateInstruction(opcode byte, operation operationFn) Instruction {
	return Instruction{
		Operand:   ImmediateAddress,
		Operation: operation,
		Length:    2,
	}
}

func ZeroPageInstruction(opcode byte, operation operationFn) Instruction {
	return Instruction{
		Operand:   ZeroPageAddress,
		Operation: operation,
		Length:    2,
	}
}

func ZeroPageXInstruction(opcode byte, operation operationFn) Instruction {
	return Instruction{
		Operand:   ZeroPageXAddress,
		Operation: operation,
		Length:    2,
	}
}

func ZeroPageYInstruction(opcode byte, operation operationFn) Instruction {
	return Instruction{
		Operand:   ZeroPageYAddress,
		Operation: operation,
		Length:    2,
	}
}

func AbsoluteInstruction(opcode byte, operation operationFn) Instruction {
	return Instruction{
		Operand:   AddressOperand,
		Operation: operation,
		Length:    3,
	}
}

func AbsoluteXInstruction(opcode byte, operation operationFn) Instruction {
	return Instruction{
		Operand:   AbsoluteXAddress,
		Operation: operation,
		Length:    3,
	}
}

func AbsoluteYInstruction(opcode byte, operation operationFn) Instruction {
	return Instruction{
		Operand:   AbsoluteYAddress,
		Operation: operation,
		Length:    3,
	}
}

func NoOperandInstruction(opcode byte, operation operationFn) Instruction {
	return Instruction{
		Operand:   NoAddress,
		Operation: operation,
		Length:    1,
	}
}

// Ops6502 is a 6502 opcode map
var Ops6502 = map[byte]Instruction{
	// complete! (except BRK)
	0x00: NoOperandInstruction(0x00, BRK),
	0x01: Instruction{IndexedIndirectAddress, ORA, 2},
	0x05: ZeroPageInstruction(0x05, ORA),
	0x06: ZeroPageInstruction(0x06, ASL),
	0x08: NoOperandInstruction(0x08, PHP),
	0x09: ImmediateInstruction(0x09, ORA),
	0x0a: NoOperandInstruction(0x0a, ASLA),
	0x0d: AbsoluteInstruction(0x0d, ORA),
	0x0e: AbsoluteInstruction(0x0e, ASL),

	// complete!
	0x10: ImmediateInstruction(0x10, BPL),
	0x11: Instruction{IndirectIndexedAddress, ORA, 2},
	0x15: ZeroPageXInstruction(0x15, ORA),
	0x16: ZeroPageXInstruction(0x16, ASL),
	0x18: NoOperandInstruction(0x18, CLC),
	0x19: AbsoluteYInstruction(0x19, ORA),
	0x1d: AbsoluteXInstruction(0x1d, ORA),
	0x1e: AbsoluteXInstruction(0x1e, ASL),

	// complete!
	0x20: AbsoluteInstruction(0x20, JSR),
	0x21: Instruction{IndexedIndirectAddress, AND, 2},
	0x24: ZeroPageInstruction(0x24, BIT),
	0x25: ZeroPageInstruction(0x25, AND),
	0x26: ZeroPageInstruction(0x26, ROL),
	0x28: NoOperandInstruction(0x28, PLP),
	0x29: ImmediateInstruction(0x29, AND),
	0x2a: NoOperandInstruction(0x2a, ROLA),
	0x2c: AbsoluteInstruction(0x2c, BIT),
	0x2d: AbsoluteInstruction(0x2d, AND),
	0x2e: AbsoluteInstruction(0x2e, ROL),

	// complete!
	0x30: ImmediateInstruction(0x30, BMI),
	0x31: Instruction{IndirectIndexedAddress, AND, 2},
	0x35: ZeroPageXInstruction(0x35, AND),
	0x36: ZeroPageXInstruction(0x36, ROL),
	0x38: NoOperandInstruction(0x38, SEC),
	0x39: AbsoluteYInstruction(0x39, AND),
	0x3d: AbsoluteXInstruction(0x3d, AND),
	0x3e: AbsoluteXInstruction(0x3e, ROL),

	// complete!
	0x40: NoOperandInstruction(0x40, RTI),
	0x41: Instruction{IndexedIndirectAddress, EOR, 2},
	0x45: ZeroPageInstruction(0x45, EOR),
	0x46: ZeroPageInstruction(0x46, LSR),
	0x48: NoOperandInstruction(0x48, PHA),
	0x49: ImmediateInstruction(0x49, EOR),
	0x4a: NoOperandInstruction(0x4a, LSRA),
	0x4c: AbsoluteInstruction(0x4c, JMP),
	0x4d: AbsoluteInstruction(0x4d, EOR),
	0x4e: AbsoluteInstruction(0x4e, LSR),

	// complete!
	0x50: ImmediateInstruction(0x50, BVC),
	0x51: Instruction{IndirectIndexedAddress, EOR, 2},
	0x55: ZeroPageXInstruction(0x55, EOR),
	0x56: ZeroPageXInstruction(0x56, LSR),
	0x58: NoOperandInstruction(0x58, CLI),
	0x59: AbsoluteYInstruction(0x59, EOR),
	0x5d: AbsoluteXInstruction(0x5d, EOR),
	0x5e: AbsoluteXInstruction(0x5e, LSR),

	// complete!
	0x60: NoOperandInstruction(0x60, RTS),
	0x61: Instruction{IndexedIndirectAddress, ADC, 2},
	0x65: ZeroPageInstruction(0x65, ADC),
	0x66: ZeroPageInstruction(0x66, ROR),
	0x68: NoOperandInstruction(0x68, PLA),
	0x69: ImmediateInstruction(0x69, ADC),
	0x6a: NoOperandInstruction(0x6a, RORA),
	0x6c: Instruction{IndirectAddress, JMP, 3},
	0x6d: AbsoluteInstruction(0x6d, ADC),
	0x6e: AbsoluteInstruction(0x6e, ROR),

	// complete!
	0x70: ImmediateInstruction(0x70, BVS),
	0x71: Instruction{IndirectIndexedAddress, ADC, 2},
	0x75: ZeroPageXInstruction(0x75, ADC),
	0x76: ZeroPageXInstruction(0x76, ROR),
	0x78: NoOperandInstruction(0x78, SEI),
	0x79: AbsoluteYInstruction(0x7d, ADC),
	0x7d: AbsoluteXInstruction(0x7d, ADC),
	0x7e: AbsoluteXInstruction(0x7e, ROR),

	// complete!
	0x81: Instruction{IndexedIndirectAddress, STA, 2},
	0x84: ZeroPageInstruction(0x84, STY),
	0x85: ZeroPageInstruction(0x85, STA),
	0x86: ZeroPageInstruction(0x86, STX),
	0x88: NoOperandInstruction(0x88, DEY),
	0x8a: NoOperandInstruction(0x8a, TXA),
	0x8c: AbsoluteInstruction(0x8c, STY),
	0x8d: AbsoluteInstruction(0x8d, STA),
	0x8e: AbsoluteInstruction(0x8e, STX),

	// complete!
	0x90: ImmediateInstruction(0x90, BCC),
	0x91: Instruction{IndirectIndexedAddress, STA, 2},
	0x94: ZeroPageXInstruction(0x94, STY),
	0x95: ZeroPageXInstruction(0x95, STA),
	0x96: ZeroPageYInstruction(0x96, STX),
	0x98: NoOperandInstruction(0x98, TYA),
	0x99: AbsoluteYInstruction(0x99, STA),
	0x9a: NoOperandInstruction(0x9a, TXS),
	0x9d: AbsoluteXInstruction(0x9d, STA),

	// complete!
	0xa0: ImmediateInstruction(0xa0, LDY),
	0xa1: Instruction{IndexedIndirectAddress, LDA, 2},
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

	// complete!
	0xb0: ImmediateInstruction(0xb0, BCS),
	0xb1: Instruction{IndirectIndexedAddress, LDA, 2},
	0xb4: ZeroPageXInstruction(0xb4, LDY),
	0xb5: ZeroPageXInstruction(0xb5, LDA),
	0xb6: ZeroPageYInstruction(0xb6, LDX),
	0xb8: NoOperandInstruction(0xb8, CLV),
	0xb9: AbsoluteYInstruction(0xb9, LDA),
	0xba: NoOperandInstruction(0xba, TSX),
	0xbc: AbsoluteXInstruction(0xbc, LDY),
	0xbd: AbsoluteXInstruction(0xbd, LDA),
	0xbe: AbsoluteYInstruction(0xbe, LDX),

	// complete!
	0xc0: ImmediateInstruction(0xc0, CPY),
	0xc1: Instruction{IndexedIndirectAddress, CMP, 2},
	0xc4: ZeroPageInstruction(0xc4, CPY),
	0xc5: ZeroPageInstruction(0xc5, CMP),
	0xc6: ZeroPageInstruction(0xc6, DEC),
	0xc8: NoOperandInstruction(0xc8, INY),
	0xc9: ImmediateInstruction(0xc9, CMP),
	0xca: NoOperandInstruction(0xca, DEX),
	0xcc: AbsoluteInstruction(0xcc, CPY),
	0xcd: AbsoluteInstruction(0xcd, CMP),
	0xce: AbsoluteInstruction(0xce, DEC),

	// complete!
	0xd0: ImmediateInstruction(0xd0, BNE),
	0xd1: Instruction{IndirectIndexedAddress, CMP, 2},
	0xd5: ZeroPageXInstruction(0xd5, CMP),
	0xd6: ZeroPageXInstruction(0xd6, DEC),
	0xd8: NoOperandInstruction(0xd8, CLD),
	0xd9: AbsoluteYInstruction(0xd9, CMP),
	0xdd: AbsoluteXInstruction(0xdd, CMP),
	0xde: AbsoluteXInstruction(0xde, DEC),

	0xe0: ImmediateInstruction(0xe0, CPX),
	0xe1: Instruction{IndexedIndirectAddress, SBC, 2},
	0xe4: ZeroPageInstruction(0xe4, CPX),
	0xe5: ZeroPageInstruction(0xe5, SBC),
	0xe6: ZeroPageInstruction(0xe6, INC),
	0xe8: NoOperandInstruction(0xe8, INX),
	0xe9: ImmediateInstruction(0xe9, SBC),
	0xea: NoOperandInstruction(0xea, NOP),
	0xec: AbsoluteInstruction(0xec, CPX),
	0xed: AbsoluteInstruction(0xed, SBC),
	0xee: AbsoluteInstruction(0xee, INC),

	0xf0: ImmediateInstruction(0xf0, BEQ),
	0xf1: Instruction{IndirectIndexedAddress, SBC, 2},
	0xf5: ZeroPageXInstruction(0xf5, SBC),
	0xf6: ZeroPageXInstruction(0xf6, INC),
	0xf8: NoOperandInstruction(0xf8, SED),
	0xf9: AbsoluteYInstruction(0xf9, SBC),
	0xfd: AbsoluteXInstruction(0xfd, SBC),
	0xfe: AbsoluteXInstruction(0xfe, INC),
}

var Ops6502Flat = flattenInstructions(Ops6502)

func flattenInstructions(input map[byte]Instruction) []*Instruction {
	output := make([]*Instruction, 256)
	for opcode, instruction := range input {
		insValue := instruction
		output[opcode] = &insValue
	}
	return output
}
