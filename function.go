package main

import (
	"fmt"
	"os"

	"github.com/DanielleB-R/goto6502/lib/cpu"
)

type Registers struct {
	A byte
	X byte
	Y byte
}

func (r *Registers) Matches(processor *cpu.Processor) bool {
	return (r.A == processor.A) && (r.X == processor.X) && (r.Y == processor.Y)
}

type MemoryMatch struct {
	base int
	data []byte
	name string
}

func (m *MemoryMatch) Matches(processor *cpu.Processor) bool {
	for offset, n := range m.data {
		if v := processor.Memory.Read(m.base + offset); v != n {
			fmt.Printf("Failure in %s, %x should be %x\n", m.name, v, n)
			return false
		}
	}
	return true
}

type Program struct {
	MachineCodeFile string
	Description     string
	FinalState      Registers
	FinalMemory     []MemoryMatch
}

func (p *Program) Check() bool {
	infile, err := os.Open(p.MachineCodeFile)
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	cpu := cpu.NewProcessor(0x1000, infile)

	for cpu.Memory.Read(cpu.PC) != 0 {
		err := cpu.Emulate()
		if err != nil {
			panic(err)
		}
	}

	for _, mem := range p.FinalMemory {
		if !mem.Matches(&cpu) {
			return false
		}
	}
	ok := p.FinalState.Matches(&cpu)
	if !ok {
		fmt.Printf("A %02x X %02x Y %02x\n", cpu.A, cpu.X, cpu.Y)
	}
	return ok
}
