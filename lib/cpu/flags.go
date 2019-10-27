package cpu

type Flags struct {
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
