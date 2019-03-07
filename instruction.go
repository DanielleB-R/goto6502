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

// ImmediateOperand gets the immediate byte operand of the current instruction
func ImmediateOperand(p *Processor) byte {
	return p.memory[p.pc+1]
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

	0xa0: &ByteInstruction{0xa0, ImmediateOperand, LDY, 2},
	0xa2: &ByteInstruction{0xa2, ImmediateOperand, LDX, 2},
	0xa4: &ByteInstruction{0xa4, ZeroPageOperand, LDY, 2},
	0xa5: &ByteInstruction{0xa5, ZeroPageOperand, LDA, 2},
	0xa6: &ByteInstruction{0xa6, ZeroPageOperand, LDX, 2},
	0xa8: &ByteInstruction{0xa8, AccumulatorOperand, LDY, 1}, // TAY
	0xa9: &ByteInstruction{0xa9, ImmediateOperand, LDA, 2},
	0xaa: &ByteInstruction{0xaa, AccumulatorOperand, LDX, 1}, // TAX
	0xac: &ByteInstruction{0xac, AbsoluteOperand, LDY, 3},
	0xad: &ByteInstruction{0xad, AbsoluteOperand, LDA, 3},
	0xae: &ByteInstruction{0xae, AbsoluteOperand, LDX, 3},

	0xb4: &ByteInstruction{0xb4, ZeroPageXOperand, LDY, 2},
	0xb5: &ByteInstruction{0xb5, ZeroPageXOperand, LDA, 2},
	0xb6: &ByteInstruction{0xb6, ZeroPageYOperand, LDX, 2},
	0xb9: &ByteInstruction{0xb9, AbsoluteYOperand, LDA, 3},
	0xbc: &ByteInstruction{0xbc, AbsoluteXOperand, LDY, 3},
	0xbd: &ByteInstruction{0xbd, AbsoluteXOperand, LDA, 3},
	0xbe: &ByteInstruction{0xbe, AbsoluteYOperand, LDX, 3},
}
