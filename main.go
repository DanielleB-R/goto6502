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

var testST = Program{
	MachineCodeFile: "asm/test-sta",
	Description:     "ST? All Addressing Modes",
	FinalState:      Registers{0x22, 0x10, 0x20},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x2000, []byte{0x22, 0x55, 0xbb}, "Absolute"},
		MemoryMatch{0x2010, []byte{0x22}, "Absolute, X"},
		MemoryMatch{0x2020, []byte{0x22}, "Absolute, Y"},
	},
}

var testINC = Program{
	MachineCodeFile: "asm/test-inc",
	Description:     "IN? All Addressing Modes",
	FinalState:      Registers{0x1f, 0x04, 0x8c},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x01, []byte{0xa7, 0xa8}, "INX"},
		MemoryMatch{0x03, []byte{0x89, 0x8c}, "INY"},
		MemoryMatch{0x2020, []byte{0xad, 0x26}, "INC Absolute"},
		MemoryMatch{0x05, []byte{0x21, 0x20}, "INC Zero Page"},
	},
}

func main() {
	fmt.Println(test1.Description, ":", test1.Check())
	fmt.Println(test2.Description, ":", test2.Check())
	fmt.Println(test4.Description, ":", test4.Check())
	fmt.Println(testLD.Description, ":", testLD.Check())
	fmt.Println(testST.Description, ":", testST.Check())
	fmt.Println(testINC.Description, ":", testINC.Check())
}
