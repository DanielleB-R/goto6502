package main

import (
	"errors"
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

func (p *Processor) ImmediateOperand() byte {
	return p.memory[p.pc+1]
}

func (p *Processor) Emulate() error {
	opcode := p.memory[p.pc]
	length := 1

	switch opcode {
	case 0xa0: // LDY immediate
		p.y = p.ImmediateOperand()
		length = 2
	case 0xa2: // LDX immediate
		p.x = p.ImmediateOperand()
		length = 2
	case 0xa9: // LDA immediate
		p.a = p.ImmediateOperand()
		length = 2
	default:
		return errors.New("unimplemented opcode")
	}

	p.pc += length
	return nil
}
