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

func (p *Processor) AddressOperand() int {
	return int(binary.LittleEndian.Uint16(p.memory[(p.pc + 1):]))
}

func AbsoluteOperand(p *Processor) byte {
	return p.memory[p.AddressOperand()]
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

func (p *Processor) Emulate() error {
	opcode := p.memory[p.pc]
	length := 1

	if ins, ok := Ops6502[opcode]; ok {
		ins.Execute(p)
		p.pc += ins.Length
		return nil
	}

	switch opcode {
	case 0x84: // STY zero page
		p.memory[ImmediateOperand(p)] = p.y
		length = 2
	case 0x85: // STA zero page
		p.memory[ImmediateOperand(p)] = p.a
		length = 2
	case 0x86: // STX zero page
		p.memory[ImmediateOperand(p)] = p.x
		length = 2
	case 0xa8: // TAY
		p.y = p.a
	case 0xaa: // TAX
		p.x = p.a
	default:
		fmt.Printf("Opcode not recognized 0x%x\n", opcode)
		return errors.New("unimplemented opcode")
	}

	p.pc += length
	return nil
}
