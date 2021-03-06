package cpu

import (
	"bytes"
	"testing"

	"github.com/DanielleB-R/goto6502/lib/memory"
	"github.com/stretchr/testify/require"
)

func makeOpcodeCpu(opcodes []byte) Processor {
	machineCode := bytes.NewReader(opcodes)
	return NewProcessor(0x1000, machineCode)
}

// func TestImmediateLoad(t *testing.T) {

// }

func TestDisableInterrupts(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x78})
	require.NoError(t, cpu.Emulate())

	require.True(t, cpu.f.I)
}

func TestEnableInterrupts(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x58})
	cpu.f.I = true

	require.NoError(t, cpu.Emulate())

	require.False(t, cpu.f.I)

}

func TestEnableDecimal(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xf8})
	require.NoError(t, cpu.Emulate())

	require.True(t, cpu.f.D)
}

func TestDisableDecimal(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xd8})
	cpu.f.D = true

	require.NoError(t, cpu.Emulate())

	require.False(t, cpu.f.D)
}

func TestRightShiftAbsolute(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x4e, 0x22, 0x22})
	cpu.Memory.Write(0x2222, 0x53)

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0x29), cpu.Memory.Read(0x2222))
	require.True(t, cpu.f.C)
}

func TestRightShiftAbsoluteX(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x5e, 0x22, 0x22})
	cpu.Memory.Write(0x2224, 0x53)
	cpu.X = 0x02

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0x29), cpu.Memory.Read(0x2224))
	require.True(t, cpu.f.C)
}

func TestRightShiftZeroPage(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x46, 0x22})
	cpu.Memory.Write(0x22, 0x53)

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0x29), cpu.Memory.Read(0x22))
	require.True(t, cpu.f.C)
}

func TestRightShiftZeroPageX(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x56, 0x22})
	cpu.Memory.Write(0x28, 0x53)
	cpu.X = 0x06

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0x29), cpu.Memory.Read(0x28))
	require.True(t, cpu.f.C)
}

func TestRotateLeftAbsolute(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x2e, 0x22, 0x22})
	cpu.Memory.Write(0x2222, 0x52)
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xa5), cpu.Memory.Read(0x2222))
	require.False(t, cpu.f.C)
}

func TestRotateLeftAbsoluteX(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x3e, 0x22, 0x22})
	cpu.Memory.Write(0x2233, 0x52)
	cpu.f.C = true
	cpu.X = 0x11

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xa5), cpu.Memory.Read(0x2233))
	require.False(t, cpu.f.C)
}

func TestRotateLeftZeroPage(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x26, 0x22})
	cpu.Memory.Write(0x22, 0x52)
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xa5), cpu.Memory.Read(0x22))
	require.False(t, cpu.f.C)
}

func TestRotateLeftZeroPageX(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x36, 0x22, 0x22})
	cpu.Memory.Write(0x33, 0x52)
	cpu.f.C = true
	cpu.X = 0x11

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xa5), cpu.Memory.Read(0x33))
	require.False(t, cpu.f.C)
}

func TestRotateRightAbsolute(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x6e, 0x22, 0x22})
	cpu.Memory.Write(0x2222, 0x52)
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xa9), cpu.Memory.Read(0x2222))
	require.False(t, cpu.f.C)
}

func TestRotateRightAbsoluteX(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x7e, 0x22, 0x22})
	cpu.Memory.Write(0x2224, 0x52)
	cpu.X = 0x02
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xa9), cpu.Memory.Read(0x2224))
	require.False(t, cpu.f.C)
}

func TestRotateRightZeroPage(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x66, 0x22})
	cpu.Memory.Write(0x22, 0x52)
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xa9), cpu.Memory.Read(0x22))
	require.False(t, cpu.f.C)
}

func TestRotateRightZeroPageX(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x76, 0x22})
	cpu.Memory.Write(0x28, 0x52)
	cpu.X = 0x06
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xa9), cpu.Memory.Read(0x28))
	require.False(t, cpu.f.C)
}

func TestBitAbsolute(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x2c, 0x22, 0x22})
	cpu.Memory.Write(0x2222, 0x86)
	cpu.A = 0x51

	require.NoError(t, cpu.Emulate())

	require.False(t, cpu.f.V)
	require.True(t, cpu.f.N)
	require.True(t, cpu.f.Z)
}

func TestBitZeroPage(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x24, 0x22})
	cpu.Memory.Write(0x22, 0x56)
	cpu.A = 0x51

	require.NoError(t, cpu.Emulate())

	require.True(t, cpu.f.V)
	require.False(t, cpu.f.N)
	require.False(t, cpu.f.Z)
}

func TestRTI(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0x40})
	cpu.push(0x38)
	cpu.push(0x84)
	cpu.push(0x81)
	cpu.f.Z = true
	cpu.f.V = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, 0x3884, cpu.PC)

	require.True(t, cpu.f.N)
	require.True(t, cpu.f.C)
	require.False(t, cpu.f.Z)
	require.False(t, cpu.f.V)
}

func TestSubtractImmediate(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xe9, 0x34})
	cpu.A = 0x66
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0x32), cpu.A)

	require.False(t, cpu.f.N)
	require.False(t, cpu.f.Z)
	require.True(t, cpu.f.C)
	require.False(t, cpu.f.V)
}

func TestSubtractAbsolute(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xed, 0x22, 0x22})
	cpu.A = 0x66
	cpu.Memory.Write(0x2222, 0x66)
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0x00), cpu.A)

	require.False(t, cpu.f.N)
	require.True(t, cpu.f.Z)
	require.True(t, cpu.f.C)
	require.False(t, cpu.f.V)
}

func TestSubtractAbsoluteX(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xfd, 0x22, 0x22})
	cpu.A = 0x66
	cpu.Memory.Write(0x2228, 0x88)
	cpu.f.C = true
	cpu.X = 0x06

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xde), cpu.A)

	require.True(t, cpu.f.N)
	require.False(t, cpu.f.Z)
	require.False(t, cpu.f.C)
	require.True(t, cpu.f.V)
}

func TestSubtractAbsoluteY(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xf9, 0x22, 0x22})
	cpu.A = 0x66
	cpu.Memory.Write(0x2228, 0x88)
	cpu.f.C = false
	cpu.Y = 0x06

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xdd), cpu.A)

	require.True(t, cpu.f.N)
	require.False(t, cpu.f.Z)
	require.False(t, cpu.f.C)
	require.True(t, cpu.f.V)
}

func TestSubtractZeroPage(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xe5, 0x22})
	cpu.A = 0x66
	cpu.Memory.Write(0x22, 0xff)
	cpu.f.C = false

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0x66), cpu.A)

	require.False(t, cpu.f.N)
	require.False(t, cpu.f.Z)
	require.False(t, cpu.f.C)
	require.False(t, cpu.f.V)
}

func TestSubtractZeroPageX(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xf5, 0x22})
	cpu.A = 0xff
	cpu.Memory.Write(0x28, 0xff)
	cpu.f.C = true
	cpu.X = 0x06

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0x00), cpu.A)

	require.False(t, cpu.f.N)
	require.True(t, cpu.f.Z)
	require.True(t, cpu.f.C)
	require.False(t, cpu.f.V)
}

func TestSubtractIndexedIndirect(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xe1, 0x38})
	cpu.X = 0x02
	cpu.Memory.Write(0x3a, 0x22)
	cpu.Memory.Write(0x3b, 0x22)
	cpu.A = 0xff
	cpu.Memory.Write(0x2222, 0x00)
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xff), cpu.A)

	require.True(t, cpu.f.N)
	require.False(t, cpu.f.Z)
	require.True(t, cpu.f.C)
	require.False(t, cpu.f.V)
}

func TestSubtractIndirectIndexed(t *testing.T) {
	cpu := makeOpcodeCpu([]byte{0xf1, 0x3a})
	cpu.Y = 0x02
	cpu.Memory.Write(0x3a, 0x22)
	cpu.Memory.Write(0x3b, 0x22)
	cpu.A = 0xff
	cpu.Memory.Write(0x2224, 0x0f)
	cpu.f.C = true

	require.NoError(t, cpu.Emulate())

	require.Equal(t, byte(0xf0), cpu.A)

	require.True(t, cpu.f.N)
	require.False(t, cpu.f.Z)
	require.True(t, cpu.f.C)
	require.False(t, cpu.f.V)
}

func TestDecimalSubtract(t *testing.T) {
	cpu := &Processor{
		A:      0x00,
		f:      Flags{D: true},
		Memory: &memory.RandomAccessMemory{Contents: []byte{0x00}},
	}

	SBC(cpu, 0x00)
	require.Equal(t, byte(0x99), cpu.A)
	require.False(t, cpu.f.C)
}
