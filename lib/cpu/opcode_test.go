package cpu

import (
	"bytes"
	"testing"

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
