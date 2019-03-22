package main

type Flags struct {
	Z bool
	N bool
}

func (f *Flags) SetZ(n byte) {
	f.Z = n == 0
}

func (f *Flags) SetN(n byte) {
	f.N = n&0x80 != 0
}
