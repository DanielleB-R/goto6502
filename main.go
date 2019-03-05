package main

import (
	"fmt"
)

var test1 = Program{
	MachineCodeFile: "asm/test1",
	Description:     "LD? Immediate",
	FinalState:      Registers{0x01, 0x10, 0x22},
}

var test2 = Program{
	MachineCodeFile: "asm/test2",
	Description:     "ST? Zero Page",
	FinalState:      Registers{0x01, 0x10, 0x22},
	FinalMemory:     []MemoryMatch{MemoryMatch{0xa0, []byte{0x01, 0x10, 0x22}, "ST? Zero Page"}},
}

var test4 = Program{
	MachineCodeFile: "asm/test4",
	Description:     "TAX/TAY",
	FinalState:      Registers{0x01, 0x01, 0x01},
}

var testLD = Program{
	MachineCodeFile: "asm/test-lda",
	Description:     "LD? All Addressing Modes",
	FinalState:      Registers{0x44, 0x01, 0x11},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x00, []byte{0x22, 0x01, 0x10}, "Zero Page"},
		MemoryMatch{0x03, []byte{0xa9, 0xa2, 0xa0}, "Absolute"},
		MemoryMatch{0x06, []byte{0xa2, 0xa9}, "Absolute, X"},
		MemoryMatch{0x08, []byte{0xa0, 0xa2}, "Absolute, Y"},
		MemoryMatch{0x0a, []byte{0x10, 0x01}, "Zero Page, X"},
		MemoryMatch{0x0c, []byte{0x01}, "Zero Page, Y"},
	},
}

func main() {
	fmt.Println(test1.Description, ":", test1.Check())
	fmt.Println(test2.Description, ":", test2.Check())
	fmt.Println(test4.Description, ":", test4.Check())
	fmt.Println(testLD.Description, ":", testLD.Check())
}
