package cpu

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Registers struct {
	A byte
	X byte
	Y byte
}

func matchRegisters(t *testing.T, reg *Registers, processor *Processor) {
	assert.Equal(t, reg.A, processor.A, "register A doesn't match")
	assert.Equal(t, reg.X, processor.X, "register X doesn't match")
	assert.Equal(t, reg.Y, processor.Y, "register Y doesn't match")
}

type MemoryMatch struct {
	base int
	data []byte
	name string
}

func matchMemory(t *testing.T, m MemoryMatch, processor *Processor) {
	for offset, n := range m.data {
		assert.Equalf(t, n, processor.Memory.Read(m.base+offset), "Failure in %s", m.name)
	}
}

type Program struct {
	MachineCodeFile string
	Description     string
	FinalState      Registers
	FinalMemory     []MemoryMatch
}

func (p *Program) checkProgram(t *testing.T) {
	infile, err := os.Open(p.MachineCodeFile)
	require.NoError(t, err)
	defer infile.Close()

	cpu := NewProcessor(0x1000, infile)

	for cpu.Memory.Read(cpu.PC) != 0 {
		err := cpu.Emulate()
		require.NoError(t, err)
	}

	for _, mem := range p.FinalMemory {
		matchMemory(t, mem, &cpu)
	}

	matchRegisters(t, &p.FinalState, &cpu)
}

func TestLegalOpcodes(t *testing.T) {
	infile, err := os.Open("../../asm/test-legal.bin")
	require.NoError(t, err)
	defer infile.Close()

	cpu := NewProcessor(0x1000, infile)

	for cpu.Memory.Read(cpu.PC) != 0 {
		err := cpu.Emulate()
		require.NoError(t, err)
		require.Equal(t, byte(0xea), cpu.Memory.Read(cpu.PC))
		cpu.PC++
	}
}

func TestRunPrograms(t *testing.T) {
	for _, program := range TestPrograms {
		t.Run(program.Description, func(tt *testing.T) {
			program.checkProgram(tt)
		})
	}
}
