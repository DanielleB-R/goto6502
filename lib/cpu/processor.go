package cpu

import (
	"errors"
	"fmt"
	"io"

	"github.com/DanielleB-R/goto6502/lib/memory"
)

type Processor struct {
	A      byte
	X      byte
	Y      byte
	S      byte
	f      Flags
	PC     int
	Memory memory.Memory
	jumped bool
}

func NewProcessor(initialPC int, rom io.Reader) Processor {
	return Processor{
		S:      0xff,
		PC:     initialPC,
		Memory: memory.NewMemoryMap(rom),
	}
}

func (p *Processor) branch(addr int) {
	offset := asSigned(p.Memory.Read(addr))
	p.PC += int(offset)
}

func (p *Processor) pull() byte {
	p.S++
	return p.Memory.Read(p.stackAddr())
}

func (p *Processor) push(value byte) {
	p.Memory.Write(p.stackAddr(), value)
	p.S--
}

func (p *Processor) stackAddr() int {
	return 0x0100 | int(p.S)
}

func ADC(p *Processor, addr int) {
	oldA := p.A
	var carry int
	if p.f.C {
		carry = 0x01
	}
	operand := p.Memory.Read(addr)
	if p.f.D {
		lna := p.A & 0x0f
		lno := operand & 0x0f
		lnr := lna + lno + byte(carry)
		if lnr > 0x09 {
			lnr -= 0x0a
			carry = 0x01
		} else {
			carry = 0x00
		}

		hna := p.A >> 4
		hno := operand >> 4
		hnr := hna + hno + byte(carry)
		if hnr > 0x09 {
			hnr -= 0x0a
			p.f.C = true
		} else {
			p.f.C = false
		}
		p.A = (hnr << 4) + (lnr & 0x0f)
	} else {
		sum := int(p.A) + int(p.Memory.Read(addr)) + carry
		p.A = byte(sum & 0xff)
		p.f.C = sum > 0xff
	}
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
	p.f.V = (oldA^operand)&0x80 == 0 && (oldA^p.A)&0x80 != 0
}

func AND(p *Processor, addr int) {
	p.A &= p.Memory.Read(addr)
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func ASLA(p *Processor, addr int) {
	p.f.C = p.A&0x80 != 0
	p.A <<= 1
	p.f.SetN(p.A)
	p.f.SetZ(p.A)
}

func ASL(p *Processor, addr int) {
	subject := p.Memory.Read(addr)
	p.f.C = subject&0x80 != 0
	subject <<= 1
	p.Memory.Write(addr, subject)
	p.f.SetN(subject)
	p.f.SetZ(subject)
}

func BCC(p *Processor, addr int) {
	if !p.f.C {
		p.branch(addr)
	}
}

func BCS(p *Processor, addr int) {
	if p.f.C {
		p.branch(addr)
	}
}

func BEQ(p *Processor, addr int) {
	if p.f.Z {
		p.branch(addr)
	}
}

func BIT(p *Processor, addr int) {
	operand := p.Memory.Read(addr)
	anded := p.A & operand

	p.f.Z = anded == 0
	p.f.V = operand&0x40 != 0
	p.f.N = operand&0x80 != 0
}

func BMI(p *Processor, addr int) {
	if p.f.N {
		p.branch(addr)
	}
}

func BNE(p *Processor, addr int) {
	if !p.f.Z {
		p.branch(addr)
	}
}

func BPL(p *Processor, addr int) {
	if !p.f.N {
		p.branch(addr)
	}
}

func BRK(p *Processor, addr int) {
	ret := p.PC + 2
	p.push(byte(ret >> 8))
	p.push(byte(ret & 0xff))
	PHP(p, 0)
	p.f.I = true
	JMP(p, p.Memory.ReadWord(0xFFFE))
}

func BVC(p *Processor, addr int) {
	if !p.f.V {
		p.branch(addr)
	}
}

func BVS(p *Processor, addr int) {
	if p.f.V {
		p.branch(addr)
	}
}

func CLC(p *Processor, addr int) {
	p.f.C = false
}

func CLD(p *Processor, addr int) {
	p.f.D = false
}

func CLI(p *Processor, addr int) {
	p.f.I = false
}

func CLV(p *Processor, addr int) {
	p.f.V = false
}

func CMP(p *Processor, addr int) {
	mem := p.Memory.Read(addr)
	diff := int(p.A) - int(mem)
	p.f.N = diff < 0
	p.f.Z = diff == 0
	p.f.C = diff >= 0
}

func CPX(p *Processor, addr int) {
	mem := p.Memory.Read(addr)
	diff := int(p.X) - int(mem)
	p.f.N = diff < 0
	p.f.Z = diff == 0
	p.f.C = diff >= 0
}

func CPY(p *Processor, addr int) {
	mem := p.Memory.Read(addr)
	diff := int(p.Y) - int(mem)
	p.f.N = diff < 0
	p.f.Z = diff == 0
	p.f.C = diff >= 0
}

func DEC(p *Processor, addr int) {
	n := p.Memory.Read(addr)
	n--
	p.Memory.Write(addr, n)
	p.f.SetZ(n)
	p.f.SetN(n)
}

func DEX(p *Processor, addr int) {
	p.X--
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

func DEY(p *Processor, addr int) {
	p.Y--
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func EOR(p *Processor, addr int) {
	p.A ^= p.Memory.Read(addr)
	p.f.SetZ(p.A)
	p.f.SetN(p.A)

}

func INC(p *Processor, addr int) {
	n := p.Memory.Read(addr)
	n++
	p.Memory.Write(addr, n)
	p.f.SetZ(n)
	p.f.SetN(n)
}

func INX(p *Processor, addr int) {
	p.X++
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

func INY(p *Processor, addr int) {
	p.Y++
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func JMP(p *Processor, addr int) {
	p.PC = addr
	p.jumped = true
}

func JSR(p *Processor, addr int) {
	ret := p.PC + 2
	p.push(byte(ret >> 8))
	p.push(byte(ret & 0xff))
	JMP(p, addr)
}

// LDA loads a byte into the A register
func LDA(p *Processor, addr int) {
	p.A = p.Memory.Read(addr)
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

// LDX loads a byte into the X register
func LDX(p *Processor, addr int) {
	p.X = p.Memory.Read(addr)
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

// LDY loads a byte into the Y register
func LDY(p *Processor, addr int) {
	p.Y = p.Memory.Read(addr)
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func LSR(p *Processor, addr int) {
	subject := p.Memory.Read(addr)
	p.f.C = subject&0x01 != 0
	subject >>= 1
	p.Memory.Write(addr, subject)
	p.f.SetZ(subject)
	p.f.SetN(subject)
}

func LSRA(p *Processor, addr int) {
	p.f.C = p.A&0x01 != 0
	p.A >>= 1
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func NOP(p *Processor, addr int) {
}

func ORA(p *Processor, addr int) {
	p.A |= p.Memory.Read(addr)
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func PHA(p *Processor, addr int) {
	p.push(p.A)
}

func PHP(p *Processor, addr int) {
	// PHP always sets bits 4 and 5
	p.push(p.f.GetByte() | 0x30)
}

func PLA(p *Processor, addr int) {
	p.A = p.pull()
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func PLP(p *Processor, addr int) {
	p.f.SetByte(p.pull())
}

func ROL(p *Processor, addr int) {
	var carry byte
	if p.f.C {
		carry = 0x01
	}

	subject := p.Memory.Read(addr)
	p.f.C = subject&0x80 != 0
	subject <<= 1
	subject |= carry
	p.Memory.Write(addr, subject)
	p.f.SetZ(subject)
	p.f.SetN(subject)
}

func ROLA(p *Processor, addr int) {
	var carry byte
	if p.f.C {
		carry = 0x01
	}
	p.f.C = p.A&0x80 != 0
	p.A <<= 1
	p.A |= carry
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func ROR(p *Processor, addr int) {
	var carry byte
	if p.f.C {
		carry = 0x80
	}

	subject := p.Memory.Read(addr)
	p.f.C = subject&0x01 != 0
	subject >>= 1
	subject |= carry
	p.Memory.Write(addr, subject)
	p.f.SetZ(subject)
	p.f.SetN(subject)
}

func RORA(p *Processor, addr int) {
	var carry byte
	if p.f.C {
		carry = 0x80
	}
	p.f.C = p.A&0x01 != 0
	p.A >>= 1
	p.A |= carry
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func RTI(p *Processor, addr int) {
	p.f.SetByte(p.pull())
	p.jumped = true
	RTS(p, addr)
}

func RTS(p *Processor, addr int) {
	lob := p.pull()
	hob := p.pull()
	p.PC = (int(hob) << 8) | int(lob)
}

func SBC(p *Processor, addr int) {
	oldA := p.A
	var carry uint
	if !p.f.C {
		carry = 0x01
	}
	operand := p.Memory.Read(addr)
	if p.f.D {
		lna := p.A & 0x0f
		lno := operand & 0x0f
		lnr := lna - lno - byte(carry)
		if lnr&0xff > 0x09 {
			lnr -= 0x06
			carry = 0x01
		} else {
			carry = 0x00
		}

		hna := p.A >> 4
		hno := operand >> 4
		hnr := hna - hno - byte(carry)
		if hnr&0xff > 0x09 {
			hnr -= 0x06
			p.f.C = false
		} else {
			p.f.C = true
		}
		p.A = (hnr << 4) + (lnr & 0x0f)
	} else {
		diff := uint(p.A) - uint(p.Memory.Read(addr)) - carry
		p.A = byte(diff & 0xff)
		p.f.C = diff <= 0xff
		p.f.SetZ(p.A)
		p.f.SetN(p.A)
		p.f.V = (oldA^operand)&0x80 != 0 && (oldA^p.A)&0x80 != 0
	}
}

func SEC(p *Processor, addr int) {
	p.f.C = true
}

func SED(p *Processor, addr int) {
	p.f.D = true
}

func SEI(p *Processor, addr int) {
	p.f.I = true
}

func STA(p *Processor, addr int) {
	p.Memory.Write(addr, p.A)
}

func STX(p *Processor, addr int) {
	p.Memory.Write(addr, p.X)
}

func STY(p *Processor, addr int) {
	p.Memory.Write(addr, p.Y)
}

func TAX(p *Processor, addr int) {
	p.X = p.A
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

func TAY(p *Processor, addr int) {
	p.Y = p.A
	p.f.SetZ(p.Y)
	p.f.SetN(p.Y)
}

func TSX(p *Processor, addr int) {
	p.X = p.S
	p.f.SetZ(p.X)
	p.f.SetN(p.X)
}

func TXA(p *Processor, addr int) {
	p.A = p.X
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func TXS(p *Processor, addr int) {
	p.S = p.X
}

func TYA(p *Processor, addr int) {
	p.A = p.Y
	p.f.SetZ(p.A)
	p.f.SetN(p.A)
}

func (p *Processor) Emulate() error {
	opcode := p.Memory.Read(p.PC)

	if ins := Ops6502Flat[opcode]; ins != nil {
		ins.Execute(p)
		if p.jumped {
			p.jumped = false
		} else {
			p.PC += ins.Length
		}
		return nil
	}

	fmt.Printf("Opcode not recognized 0x%x\n", opcode)
	return errors.New("unimplemented opcode")
}
