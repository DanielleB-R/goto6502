package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Processor struct {
	a      byte
	x      byte
	y      byte
	f      Flags
	pc     int
	memory [65536]byte
	jumped bool
}

func (p *Processor) LoadMemory(filename string, base int) error {
	infile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer infile.Close()

	_, err = infile.Read(p.memory[base:])
	if err != nil {
		return err
	}
	return nil
}

func (p *Processor) byteAt(addr int) byte {
	return p.memory[addr]
}

func (p *Processor) addressAt(addr int) int {
	return int(binary.LittleEndian.Uint16(p.memory[addr:]))
}

func (p *Processor) branch(addr int) {
	offset := asSigned(p.byteAt(addr))
	p.pc += int(offset)
}

func BEQ(p *Processor, addr int) {
	if p.f.Z {
		p.branch(addr)
	}
}

func DEC(p *Processor, addr int) {
	p.memory[addr]--
	p.f.SetZ(p.memory[addr])
	p.f.SetN(p.memory[addr])
}

func DEX(p *Processor, addr int) {
	p.x--
	p.f.SetZ(p.x)
	p.f.SetN(p.x)
}

func DEY(p *Processor, addr int) {
	p.y--
	p.f.SetZ(p.y)
	p.f.SetN(p.y)
}

func INC(p *Processor, addr int) {
	p.memory[addr]++
	p.f.SetZ(p.memory[addr])
	p.f.SetN(p.memory[addr])
}

func INX(p *Processor, addr int) {
	p.x++
	p.f.SetZ(p.x)
	p.f.SetN(p.x)
}

func INY(p *Processor, addr int) {
	p.y++
	p.f.SetZ(p.y)
	p.f.SetN(p.y)
}

func JMP(p *Processor, addr int) {
	p.pc = addr
	p.jumped = true
}

// LDA loads a byte into the A register
func LDA(p *Processor, addr int) {
	p.a = p.byteAt(addr)
	p.f.SetZ(p.a)
	p.f.SetN(p.a)
}

// LDX loads a byte into the X register
func LDX(p *Processor, addr int) {
	p.x = p.byteAt(addr)
	p.f.SetZ(p.x)
	p.f.SetN(p.x)
}

// LDY loads a byte into the Y register
func LDY(p *Processor, addr int) {
	p.y = p.byteAt(addr)
	p.f.SetZ(p.y)
	p.f.SetN(p.y)
}

func STA(p *Processor, address int) {
	p.memory[address] = p.a
}

func STX(p *Processor, address int) {
	p.memory[address] = p.x
}

func STY(p *Processor, address int) {
	p.memory[address] = p.y
}

func TAX(p *Processor, addr int) {
	p.x = p.a
	p.f.SetZ(p.x)
	p.f.SetN(p.x)
}

func TAY(p *Processor, addr int) {
	p.y = p.a
	p.f.SetZ(p.y)
	p.f.SetN(p.y)
}

func (p *Processor) Emulate() error {
	opcode := p.memory[p.pc]

	if ins, ok := Ops6502[opcode]; ok {
		ins.Execute(p)
		if p.jumped {
			p.jumped = false
		} else {
			p.pc += ins.Length
		}
		return nil
	}

	fmt.Printf("Opcode not recognized 0x%x\n", opcode)
	return errors.New("unimplemented opcode")
}
