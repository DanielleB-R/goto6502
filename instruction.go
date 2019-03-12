package main

// Instruction data

// Instruction contains the details for a single 6502 instruction
type Instruction struct {
	Opcode    byte
	Operand   func(*Processor) int
	Operation func(*Processor, int)
	Length    int
}

// Execute executes the instruction against the current state of the CPU
func (i *Instruction) Execute(cpu *Processor) {
	i.Operation(cpu, i.Operand(cpu))
}

// Operand functions

func ImmediateAddress(p *Processor) int {
	return p.pc + 1
}

// ImmediateOperand gets the immediate byte operand of the current instruction
func ImmediateOperand(p *Processor) byte {
	return p.byteAt(p.pc + 1)
}

func ZeroPageAddress(p *Processor) int {
	return int(ImmediateOperand(p))
}

func ZeroPageXAddress(p *Processor) int {
	return int(ImmediateOperand(p) + p.x)
}

func ZeroPageYAddress(p *Processor) int {
	return int(ImmediateOperand(p) + p.y)
}

// 6502 opcode map

var Ops6502 = map[byte]Instruction{
	0x84: Instruction{0x84, ZeroPageAddress, STY, 2},
	0x85: Instruction{0x85, ZeroPageAddress, STA, 2},
	0x86: Instruction{0x86, ZeroPageAddress, STX, 2},
	0x8c: Instruction{0x8c, AddressOperand, STY, 3},
	0x8d: Instruction{0x8d, AddressOperand, STA, 3},
	0x8e: Instruction{0x8e, AddressOperand, STX, 3},

	0x99: Instruction{0x99, AbsoluteYAddress, STA, 3},
	0x9d: Instruction{0x9d, AbsoluteXAddress, STA, 3},

	0xa0: Instruction{0xa0, ImmediateAddress, LDY, 2},
	0xa2: Instruction{0xa2, ImmediateAddress, LDX, 2},
	0xa4: Instruction{0xa4, ZeroPageAddress, LDY, 2},
	0xa5: Instruction{0xa5, ZeroPageAddress, LDA, 2},
	0xa6: Instruction{0xa6, ZeroPageAddress, LDX, 2},
	0xa8: Instruction{0xa8, AddressOperand, TAY, 1}, // TAY
	0xa9: Instruction{0xa9, ImmediateAddress, LDA, 2},
	0xaa: Instruction{0xaa, AddressOperand, TAX, 1}, // TAX
	0xac: Instruction{0xac, AddressOperand, LDY, 3},
	0xad: Instruction{0xad, AddressOperand, LDA, 3},
	0xae: Instruction{0xae, AddressOperand, LDX, 3},

	0xb4: Instruction{0xb4, ZeroPageXAddress, LDY, 2},
	0xb5: Instruction{0xb5, ZeroPageXAddress, LDA, 2},
	0xb6: Instruction{0xb6, ZeroPageYAddress, LDX, 2},
	0xb9: Instruction{0xb9, AbsoluteYAddress, LDA, 3},
	0xbc: Instruction{0xbc, AbsoluteXAddress, LDY, 3},
	0xbd: Instruction{0xbd, AbsoluteXAddress, LDA, 3},
	0xbe: Instruction{0xbe, AbsoluteYAddress, LDX, 3},
}
