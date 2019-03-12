package main

// Instruction data

type ByteInstruction struct {
	Opcode    byte
	Operand   func(*Processor) byte
	Operation func(*Processor, byte)
	length    int
}

// Execute executes the instruction against the current state of the CPU
func (i *ByteInstruction) Execute(cpu *Processor) {
	i.Operation(cpu, i.Operand(cpu))
}

func (i *ByteInstruction) Length() int {
	return i.length
}

type AddressInstruction struct {
	Opcode    byte
	Operand   func(*Processor) int
	Operation func(*Processor, int)
	length    int
}

// Execute executes the instruction against the current state of the CPU
func (i *AddressInstruction) Execute(cpu *Processor) {
	i.Operation(cpu, i.Operand(cpu))
}

func (i *AddressInstruction) Length() int {
	return i.length
}

// Instruction contains the details for a single 6502 instruction
type Instruction interface {
	Execute(cpu *Processor)
	Length() int
}

// Operand functions

func ImmediateAddress(p *Processor) int {
	return p.pc + 1
}

// ImmediateOperand gets the immediate byte operand of the current instruction
func ImmediateOperand(p *Processor) byte {
	return p.byteAt(p.pc + 1)
}

// ZeroPageOperand gets the zero page byte operand of the current instruction
func ZeroPageOperand(p *Processor) byte {
	return p.memory[ImmediateOperand(p)]
}

func ZeroPageXOperand(p *Processor) byte {
	// Because these values are bytes they wrap normally at 0xff
	return p.memory[ImmediateOperand(p)+p.x]
}

func ZeroPageYOperand(p *Processor) byte {
	// Because these values are bytes they wrap normally at 0xff
	return p.memory[ImmediateOperand(p)+p.y]
}

func AccumulatorOperand(p *Processor) byte {
	return p.a
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
	0x84: &AddressInstruction{0x84, ZeroPageAddress, STY, 2},
	0x85: &AddressInstruction{0x85, ZeroPageAddress, STA, 2},
	0x86: &AddressInstruction{0x86, ZeroPageAddress, STX, 2},
	0x8c: &AddressInstruction{0x8c, AddressOperand, STY, 3},
	0x8d: &AddressInstruction{0x8d, AddressOperand, STA, 3},
	0x8e: &AddressInstruction{0x8e, AddressOperand, STX, 3},

	0x99: &AddressInstruction{0x99, AbsoluteYAddress, STA, 3},
	0x9d: &AddressInstruction{0x9d, AbsoluteXAddress, STA, 3},

	0xa0: &AddressInstruction{0xa0, ImmediateAddress, LDY, 2},
	0xa2: &AddressInstruction{0xa2, ImmediateAddress, LDX, 2},
	0xa4: &AddressInstruction{0xa4, ZeroPageAddress, LDY, 2},
	0xa5: &AddressInstruction{0xa5, ZeroPageAddress, LDA, 2},
	0xa6: &AddressInstruction{0xa6, ZeroPageAddress, LDX, 2},
	0xa8: &AddressInstruction{0xa8, AddressOperand, TAY, 1}, // TAY
	0xa9: &AddressInstruction{0xa9, ImmediateAddress, LDA, 2},
	0xaa: &AddressInstruction{0xaa, AddressOperand, TAX, 1}, // TAX
	0xac: &AddressInstruction{0xac, AddressOperand, LDY, 3},
	0xad: &AddressInstruction{0xad, AddressOperand, LDA, 3},
	0xae: &AddressInstruction{0xae, AddressOperand, LDX, 3},

	0xb4: &AddressInstruction{0xb4, ZeroPageXAddress, LDY, 2},
	0xb5: &AddressInstruction{0xb5, ZeroPageXAddress, LDA, 2},
	0xb6: &AddressInstruction{0xb6, ZeroPageYAddress, LDX, 2},
	0xb9: &AddressInstruction{0xb9, AbsoluteYAddress, LDA, 3},
	0xbc: &AddressInstruction{0xbc, AbsoluteXAddress, LDY, 3},
	0xbd: &AddressInstruction{0xbd, AbsoluteXAddress, LDA, 3},
	0xbe: &AddressInstruction{0xbe, AbsoluteYAddress, LDX, 3},
}
