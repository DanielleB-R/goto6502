package main

import (
	"fmt"
)

var test1 = Program{
	MachineCodeFile: "asm/test1.bin",
	Description:     "LD? Immediate",
	FinalState:      Registers{0x01, 0x10, 0x22},
}

var test2 = Program{
	MachineCodeFile: "asm/test2.bin",
	Description:     "ST? Zero Page",
	FinalState:      Registers{0x01, 0x10, 0x22},
	FinalMemory:     []MemoryMatch{MemoryMatch{0xa0, []byte{0x01, 0x10, 0x22}, "ST? Zero Page"}},
}

var test4 = Program{
	MachineCodeFile: "asm/test4.bin",
	Description:     "TAX/TAY",
	FinalState:      Registers{0x01, 0x01, 0x01},
}

var testLD = Program{
	MachineCodeFile: "asm/test-lda.bin",
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
	MachineCodeFile: "asm/test-sta.bin",
	Description:     "ST? All Addressing Modes",
	FinalState:      Registers{0x80, 0xef, 0x22},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x2000, []byte{0x22, 0x55, 0xbb}, "Absolute"},
		MemoryMatch{0x2010, []byte{0x22}, "Absolute, X"},
		MemoryMatch{0x2020, []byte{0x22}, "Absolute, Y"},
		MemoryMatch{0x30, []byte{0x80, 0x75}, "Zero Page, X"},
		MemoryMatch{0x33, []byte{0xef}, "Zero Page, Y"},
	},
}

var testINC = Program{
	MachineCodeFile: "asm/test-inc.bin",
	Description:     "IN? All Addressing Modes",
	FinalState:      Registers{0x1f, 0x04, 0x8c},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x01, []byte{0xa7, 0xa8}, "INX"},
		MemoryMatch{0x03, []byte{0x89, 0x8c}, "INY"},
		MemoryMatch{0x2020, []byte{0xad, 0x26}, "INC Absolute"},
		MemoryMatch{0x05, []byte{0x21, 0x20}, "INC Zero Page"},
	},
}

var testDEC = Program{
	MachineCodeFile: "asm/test-dec.bin",
	Description:     "DE? All Addressing Modes",
	FinalState:      Registers{0x1f, 0x04, 0x84},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x01, []byte{0xa3, 0xa2}, "DEX"},
		MemoryMatch{0x03, []byte{0x87, 0x84}, "DEY"},
		MemoryMatch{0x2020, []byte{0xa9, 0x22}, "DEC Absolute"},
		MemoryMatch{0x05, []byte{0x1d, 0x1e}, "DEC Zero Page"},
	},
}

var testJMP = Program{
	MachineCodeFile: "asm/test-jmp.bin",
	Description:     "JMP Absolute and Indirect",
	FinalState:      Registers{0x22, 0x20, 0x99},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x01, []byte{0x05, 0x20}, "JMP Absolute"},
		MemoryMatch{0x03, []byte{0x22}, "JMP Indirect"},
	},
}

var testBranch = Program{
	MachineCodeFile: "asm/test-branch.bin",
	Description:     "Branch",
	FinalState:      Registers{0xfa, 0x9f, 0x82},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x00, []byte{0x00}, "No error"},
		MemoryMatch{0x01, []byte{0x28, 0xcd}, "BEQ"},
		MemoryMatch{0x03, []byte{0x9f, 0x82}, "BNE"},
	},
}

var testStrcpy = Program{
	MachineCodeFile: "asm/strcpy.bin",
	Description:     "Simple STRCPY",
	FinalState:      Registers{0x00, 0x12, 0x00},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x3000, []byte("Running a program!"), "Destination"},
	},
}

var testAnd = Program{
	MachineCodeFile: "asm/test-and.bin",
	Description:     "AND instructions",
	FinalState:      Registers{0x40, 0x00, 0x00},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x00, []byte{0x40}, "AND Immediate"},
	},
}

var tests = []Program{test1, test2, test4, testLD, testST, testINC, testDEC, testJMP, testBranch, testStrcpy, testAnd}

func main() {
	for _, test := range tests {
		fmt.Println(test.Description, ":", test.Check())
	}
}
