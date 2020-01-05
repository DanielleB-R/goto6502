package cpu

// Flags represents the state of the CPU status flags
type Flags struct {
	// This flag is for interrupt disable
	I bool
	Z bool
	N bool
	C bool
	V bool
}

func (f *Flags) SetZ(n byte) {
	f.Z = n == 0
}

func (f *Flags) SetN(n byte) {
	f.N = n&0x80 != 0
}

func (f *Flags) SetV(n, oldN byte) {
	f.V = n&0x80 != oldN&0x80
}

const (
	FLAG_C byte = 0x01
	FLAG_Z byte = 0x02
	FLAG_I byte = 0x04
	FLAG_V byte = 0x40
	FLAG_N byte = 0x80
)

func (f *Flags) GetByte() byte {
	var flagByte byte
	if f.Z {
		flagByte |= FLAG_Z
	}
	if f.C {
		flagByte |= FLAG_C
	}
	if f.V {
		flagByte |= FLAG_V
	}
	if f.N {
		flagByte |= FLAG_N
	}
	if f.I {
		flagByte |= FLAG_I
	}
	return flagByte
}

func (f *Flags) SetByte(value byte) {
	f.Z = value&FLAG_Z != 0
	f.C = value&FLAG_C != 0
	f.V = value&FLAG_V != 0
	f.N = value&FLAG_N != 0
	f.I = value&FLAG_I != 0
}
