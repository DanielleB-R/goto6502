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
	pc     int
	memory [65536]byte
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

func DEC(p *Processor, addr int) {
	p.memory[addr]--
}

func DEX(p *Processor, addr int) {
	p.x--
}

func DEY(p *Processor, addr int) {
	p.y--
}

func INC(p *Processor, addr int) {
	p.memory[addr]++
}

func INX(p *Processor, addr int) {
	p.x++
}

func INY(p *Processor, addr int) {
	p.y++
}

// LDA loads a byte into the A register
func LDA(p *Processor, addr int) {
	p.a = p.byteAt(addr)
}

// LDX loads a byte into the X register
func LDX(p *Processor, addr int) {
	p.x = p.byteAt(addr)
}

// LDY loads a byte into the Y register
func LDY(p *Processor, addr int) {
	p.y = p.byteAt(addr)
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
}

func TAY(p *Processor, addr int) {
	p.y = p.a
}

func (p *Processor) Emulate() error {
	opcode := p.memory[p.pc]

	if ins, ok := Ops6502[opcode]; ok {
		ins.Execute(p)
		p.pc += ins.Length
		return nil
	}

	fmt.Printf("Opcode not recognized 0x%x\n", opcode)
	return errors.New("unimplemented opcode")
}
