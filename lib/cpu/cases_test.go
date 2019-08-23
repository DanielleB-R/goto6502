package cpu

var TestPrograms = []Program{
	Program{
		MachineCodeFile: "../../asm/test1.bin",
		Description:     "LD? Immediate",
		FinalState:      Registers{0x01, 0x10, 0x22},
	},
	Program{
		MachineCodeFile: "../../asm/test2.bin",
		Description:     "ST? Zero Page",
		FinalState:      Registers{0x01, 0x10, 0x22},
		FinalMemory:     []MemoryMatch{MemoryMatch{0xa0, []byte{0x01, 0x10, 0x22}, "ST? Zero Page"}},
	},
	Program{
		MachineCodeFile: "../../asm/test4.bin",
		Description:     "TAX/TAY",
		FinalState:      Registers{0x01, 0x01, 0x01},
	},

	Program{
		MachineCodeFile: "../../asm/test-lda.bin",
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
	},

	Program{
		MachineCodeFile: "../../asm/test-sta.bin",
		Description:     "ST? All Addressing Modes",
		FinalState:      Registers{0x80, 0xef, 0x22},
		FinalMemory: []MemoryMatch{
			MemoryMatch{0x2000, []byte{0x22, 0x55, 0xbb}, "Absolute"},
			MemoryMatch{0x2010, []byte{0x22}, "Absolute, X"},
			MemoryMatch{0x2020, []byte{0x22}, "Absolute, Y"},
			MemoryMatch{0x30, []byte{0x80, 0x75}, "Zero Page, X"},
			MemoryMatch{0x33, []byte{0xef}, "Zero Page, Y"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-inc.bin",
		Description:     "IN? All Addressing Modes",
		FinalState:      Registers{0x1f, 0x04, 0x8c},
		FinalMemory: []MemoryMatch{
			MemoryMatch{0x01, []byte{0xa7, 0xa8}, "INX"},
			MemoryMatch{0x03, []byte{0x89, 0x8c}, "INY"},
			MemoryMatch{0x2020, []byte{0xad, 0x26}, "INC Absolute"},
			MemoryMatch{0x05, []byte{0x21, 0x20}, "INC Zero Page"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-dec.bin",
		Description:     "DE? All Addressing Modes",
		FinalState:      Registers{0x1f, 0x04, 0x84},
		FinalMemory: []MemoryMatch{
			MemoryMatch{0x01, []byte{0xa3, 0xa2}, "DEX"},
			MemoryMatch{0x03, []byte{0x87, 0x84}, "DEY"},
			MemoryMatch{0x2020, []byte{0xa9, 0x22}, "DEC Absolute"},
			MemoryMatch{0x05, []byte{0x1d, 0x1e}, "DEC Zero Page"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-jmp.bin",
		Description:     "JMP Absolute and Indirect",
		FinalState:      Registers{0x22, 0x20, 0x99},
		FinalMemory: []MemoryMatch{
			MemoryMatch{0x01, []byte{0x05, 0x20}, "JMP Absolute"},
			MemoryMatch{0x03, []byte{0x22}, "JMP Indirect"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-branch.bin",
		Description:     "Branch",
		FinalState:      Registers{0xfa, 0x9f, 0x82},
		FinalMemory: []MemoryMatch{
			MemoryMatch{0x00, []byte{0x00}, "No error"},
			MemoryMatch{0x01, []byte{0x28, 0xcd}, "BEQ"},
			MemoryMatch{0x03, []byte{0x9f, 0x82}, "BNE"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/strcpy.bin",
		Description:     "Simple STRCPY",
		FinalState:      Registers{0x00, 0x12, 0x00},
		FinalMemory: []MemoryMatch{
			MemoryMatch{0x3000, []byte("Running a program!"), "Destination"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-and.bin",
		Description:     "AND instructions",
		FinalState:      Registers{0x81, 0x02, 0x03},
		FinalMemory: []MemoryMatch{
			MemoryMatch{0x00, []byte{0x40}, "AND Immediate"},
			MemoryMatch{0x01, []byte{0x90}, "AND Absolute"},
			MemoryMatch{0x02, []byte{0x09}, "AND Absolute X"},
			MemoryMatch{0x03, []byte{0x81}, "AND Absolute Y"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-ora.bin",
		Description:     "ORA instructions",
		FinalState:      Registers{0x9d, 0x02, 0x03},
		FinalMemory: []MemoryMatch{
			MemoryMatch{0x00, []byte{0xf4}, "ORA Immediate"},
			MemoryMatch{0x01, []byte{0xf9}, "ORA Absolute"},
			MemoryMatch{0x02, []byte{0x9f}, "ORA Absolute X"},
			MemoryMatch{0x03, []byte{0x9d}, "ORA Absolute Y"},
		},
	},
}
