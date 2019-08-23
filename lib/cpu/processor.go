package cpu

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Processor struct {
	A      byte
	X      byte
	Y      byte
	f      Flags
	PC     int
	Memory [65536]byte
	jumped bool
}

func (p *Processor) LoadMemory(filename string, base int) error {
	infile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer infile.Close()

	_, err = infile.Read(p.Memory[base:])
	if err != nil {
		return err
	}
	return nil
}

func (p *Processor) byteAt(addr int) byte {
	return p.Memory[addr]
}

func (p *Processor) addressAt(addr int) int {
	return int(binary.LittleEndian.Uint16(p.Memory[addr:]))
}

func (p *Processor) branch(addr int) {
	offset := asSigned(p.byteAt(addr))
	p.PC += int(offset)
}

func AND(p *Processor, addr int) {
	p.A &= p.Memory[addr]
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
	p.Memory[addr]--
	p.f.SetZ(p.Memory[addr])
	p.f.SetN(p.Memory[addr])
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
	p.Memory[addr]++
	p.f.SetZ(p.Memory[addr])
	p.f.SetN(p.Memory[addr])
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
	p.A = p.byteAt(addr)
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

// LDX loads a byte into the X register
func LDX(p *Processor, addr int) {
	p.X = p.byteAt(addr)
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

// LDY loads a byte into the Y register
func LDY(p *Processor, addr int) {
	p.Y = p.byteAt(addr)
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func ORA(p *Processor, addr int) {
	p.A |= p.Memory[addr]
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func STA(p *Processor, address int) {
	p.Memory[address] = p.A
}

func STX(p *Processor, address int) {
	p.Memory[address] = p.X
}

func STY(p *Processor, address int) {
	p.Memory[address] = p.Y
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
	opcode := p.Memory[p.PC]

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
