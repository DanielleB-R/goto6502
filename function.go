package main

type Registers struct {
	A byte
	X byte
	Y byte
}

func (r *Registers) Matches(cpu *Processor) bool {
	return (r.A == cpu.a) && (r.X == cpu.x) && (r.Y == cpu.y)
}

type MemoryMatch struct {
	base int
	data []byte
}

func (m *MemoryMatch) Matches(cpu *Processor) bool {
	for offset, n := range m.data {
		if cpu.memory[m.base+offset] != n {
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
	cpu := Processor{pc: 0x1000}
	cpu.LoadMemory(p.MachineCodeFile, 0x1000)

	for cpu.memory[cpu.pc] != 0 {
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
	return p.FinalState.Matches(&cpu)
}
