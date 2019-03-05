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

func (p *Processor) addressAt(addr int) int {
	return int(binary.LittleEndian.Uint16(p.memory[addr:]))
}

func AddressOperand(p *Processor) int {
	return p.addressAt(p.pc + 1)
}

func AbsoluteOperand(p *Processor) byte {
	return p.memory[AddressOperand(p)]
}

func AbsoluteXOperand(p *Processor) byte {
	return p.memory[AddressOperand(p)+int(p.x)]
}

func AbsoluteYOperand(p *Processor) byte {
	return p.memory[AddressOperand(p)+int(p.y)]
}

func IndexedIndirectOperand(p *Processor) byte {
	return p.memory[p.addressAt(int(ImmediateOperand(p)+p.x))]
}

// LDA loads a byte into the A register
func LDA(p *Processor, operand byte) {
	p.a = operand
}

// LDX loads a byte into the X register
func LDX(p *Processor, operand byte) {
	p.x = operand
}

// LDY loads a byte into the Y register
func LDY(p *Processor, operand byte) {
	p.y = operand
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

func (p *Processor) Emulate() error {
	opcode := p.memory[p.pc]

	if ins, ok := Ops6502[opcode]; ok {
		ins.Execute(p)
		p.pc += ins.Length()
		return nil
	}

	fmt.Printf("Opcode not recognized 0x%x\n", opcode)
	return errors.New("unimplemented opcode")
}
