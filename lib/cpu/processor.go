package cpu

import (
	"errors"
	"fmt"
	"io"

	"github.com/DanielleB-R/goto6502/lib/memory"
)

type Processor struct {
	A      byte
	X      byte
	Y      byte
	f      Flags
	PC     int
	Memory memory.Memory
	jumped bool
}

func NewProcessor(initialPC int, rom io.Reader) Processor {
	return Processor{
		PC:     initialPC,
		Memory: memory.NewMemoryMap(rom),
	}
}

func (p *Processor) branch(addr int) {
	offset := asSigned(p.Memory.Read(addr))
	p.PC += int(offset)
}

func AND(p *Processor, addr int) {
	p.A &= p.Memory.Read(addr)
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func BEQ(p *Processor, addr int) {
	if p.f.Z {
		p.branch(addr)
	}
}

func BNE(p *Processor, addr int) {
	if !p.f.Z {
		p.branch(addr)
	}
}

func DEC(p *Processor, addr int) {
	n := p.Memory.Read(addr)
	n--
	p.Memory.Write(addr, n)
	p.f.SetZ(n)
	p.f.SetN(n)
}

func DEX(p *Processor, addr int) {
	p.X--
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

func DEY(p *Processor, addr int) {
	p.Y--
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func INC(p *Processor, addr int) {
	n := p.Memory.Read(addr)
	n++
	p.Memory.Write(addr, n)
	p.f.SetZ(n)
	p.f.SetN(n)
}

func INX(p *Processor, addr int) {
	p.X++
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

func INY(p *Processor, addr int) {
	p.Y++
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func JMP(p *Processor, addr int) {
	p.PC = addr
	p.jumped = true
}

// LDA loads a byte into the A register
func LDA(p *Processor, addr int) {
	p.A = p.Memory.Read(addr)
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

// LDX loads a byte into the X register
func LDX(p *Processor, addr int) {
	p.X = p.Memory.Read(addr)
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

// LDY loads a byte into the Y register
func LDY(p *Processor, addr int) {
	p.Y = p.Memory.Read(addr)
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func ORA(p *Processor, addr int) {
	p.A |= p.Memory.Read(addr)
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func STA(p *Processor, address int) {
	p.Memory.Write(address, p.A)
}

func STX(p *Processor, address int) {
	p.Memory.Write(address, p.X)
}

func STY(p *Processor, address int) {
	p.Memory.Write(address, p.Y)
}

func TAX(p *Processor, addr int) {
	p.X = p.A
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

func TAY(p *Processor, addr int) {
	p.Y = p.A
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func (p *Processor) Emulate() error {
	opcode := p.Memory.Read(p.PC)

	if ins, ok := Ops6502[opcode]; ok {
		ins.Execute(p)
		if p.jumped {
			p.jumped = false
		} else {
			p.PC += ins.Length
		}
		return nil
	}

	fmt.Printf("Opcode not recognized 0x%x\n", opcode)
	return errors.New("unimplemented opcode")
}
