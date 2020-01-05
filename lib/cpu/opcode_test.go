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
