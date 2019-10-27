package cpu

var TestPrograms = []Program{
	Program{
		MachineCodeFile: "../../asm/load-immediate.bin",
		Description:     "LD? Immediate",
		FinalState:      Registers{0x01, 0x10, 0x22},
	},
	Program{
		MachineCodeFile: "../../asm/test2.bin",
		Description:     "ST? Zero Page",
		FinalState:      Registers{0x01, 0x10, 0x22},
		FinalMemory:     []MemoryMatch{{0xa0, []byte{0x01, 0x10, 0x22}, "ST? Zero Page"}},
	},
	Program{
		MachineCodeFile: "../../asm/test4.bin",
		Description:     "TAX/TAY",
		FinalState:      Registers{0x01, 0x01, 0x01},
	},

	Program{
		MachineCodeFile: "../../asm/test-lda.bin",
		Description:     "LD? All Addressing Modes",
		FinalState:      Registers{0xa2, 0x01, 0x02},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0x22, 0x01, 0x10}, "Zero Page"},
			{0x03, []byte{0xa9, 0xa2, 0xa0}, "Absolute"},
			{0x06, []byte{0xa2, 0xa9}, "Absolute, X"},
			{0x08, []byte{0xa0, 0xa2}, "Absolute, Y"},
			{0x0a, []byte{0x10, 0x01}, "Zero Page, X"},
			{0x0c, []byte{0x01}, "Zero Page, Y"},
			{0x0d, []byte{0xa9}, "Indexed Indirect, X"},
			{0x0e, []byte{0xa2}, "Indirect Indexed, Y"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-sta.bin",
		Description:     "ST? All Addressing Modes",
		FinalState:      Registers{0xdd, 0x01, 0x03},
		FinalMemory: []MemoryMatch{
			{0x2000, []byte{0x22, 0x55, 0xbb}, "Absolute"},
			{0x2010, []byte{0x22}, "Absolute, X"},
			{0x2020, []byte{0x22}, "Absolute, Y"},
			{0x30, []byte{0x80, 0x75}, "Zero Page, X"},
			{0x33, []byte{0xef}, "Zero Page, Y"},
			{0x35, []byte{0xce}, "Indexed Indirect, X"},
			{0x38, []byte{0xdd}, "Indirect Indexed, Y"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-inc.bin",
		Description:     "IN? All Addressing Modes",
		FinalState:      Registers{0x1f, 0x04, 0x8c},
		FinalMemory: []MemoryMatch{
			{0x01, []byte{0xa7, 0xa8}, "INX"},
			{0x03, []byte{0x89, 0x8c}, "INY"},
			{0x2020, []byte{0xad, 0x26}, "INC Absolute"},
			{0x05, []byte{0x21, 0x20}, "INC Zero Page"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-dec.bin",
		Description:     "DE? All Addressing Modes",
		FinalState:      Registers{0x1f, 0x04, 0x84},
		FinalMemory: []MemoryMatch{
			{0x01, []byte{0xa3, 0xa2}, "DEX"},
			{0x03, []byte{0x87, 0x84}, "DEY"},
			{0x2020, []byte{0xa9, 0x22}, "DEC Absolute"},
			{0x05, []byte{0x1d, 0x1e}, "DEC Zero Page"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-jmp.bin",
		Description:     "JMP Absolute and Indirect",
		FinalState:      Registers{0x22, 0x20, 0x99},
		FinalMemory: []MemoryMatch{
			{0x01, []byte{0x05, 0x20}, "JMP Absolute"},
			{0x03, []byte{0x22}, "JMP Indirect"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-branch.bin",
		Description:     "Branch",
		FinalState:      Registers{0xf9, 0xde, 0x8b},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0x00}, "No error"},
			{0x01, []byte{0x28, 0xcd}, "BEQ"},
			{0x03, []byte{0x9f, 0x82}, "BNE"},
			{0x05, []byte{0xde, 0x8b}, "BPL"},
			{0x07, []byte{0xde, 0x8b}, "BMI"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/strcpy.bin",
		Description:     "Simple STRCPY",
		FinalState:      Registers{0x00, 0x12, 0x00},
		FinalMemory: []MemoryMatch{
			{0x3000, []byte("Running a program!"), "Destination"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-and.bin",
		Description:     "AND instructions",
		FinalState:      Registers{0x10, 0x0c, 0x15},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0x40}, "AND Immediate"},
			{0x01, []byte{0x90}, "AND Absolute"},
			{0x02, []byte{0x09}, "AND Absolute X"},
			{0x03, []byte{0x81}, "AND Absolute Y"},
			{0x04, []byte{0xc1}, "AND Indexed Indirect"},
			{0x05, []byte{0x99}, "AND Indirect Indexed"},
			{0x06, []byte{0x01}, "AND Zero Page"},
			{0x07, []byte{0x10}, "AND Zero Page"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-ora.bin",
		Description:     "ORA instructions",
		FinalState:      Registers{0xdd, 0x11, 0x15},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0xf4}, "ORA Immediate"},
			{0x01, []byte{0xf9}, "ORA Absolute"},
			{0x02, []byte{0x9f}, "ORA Absolute X"},
			{0x03, []byte{0x9d}, "ORA Absolute Y"},
			{0x04, []byte{0xd9}, "ORA Zero Page"},
			{0x05, []byte{0x11}, "ORA Zero Page X"},
			{0x06, []byte{0xf5}, "ORA Indexed Indirect"},
			{0x07, []byte{0xdd}, "ORA Indirect Indexed"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-eor.bin",
		Description:     "EOR instructions",
		FinalState:      Registers{0x22, 0x11, 0x15},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0xb4}, "EOR Immediate"},
			{0x01, []byte{0x69}, "EOR Absolute"},
			{0x02, []byte{0x96}, "EOR Absolute X"},
			{0x03, []byte{0x1c}, "EOR Absolute Y"},
			{0x04, []byte{0xca}, "EOR Zero Page"},
			{0x05, []byte{0x68}, "EOR Zero Page X"},
			{0x06, []byte{0x34}, "EOR Indexed Indirect"},
			{0x07, []byte{0x22}, "EOR Indirect Indexed"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-carry.bin",
		Description:     "Carry flag instructions",
		FinalState:      Registers{0x27, 0x00, 0x00},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0x00}, "No error"},
			{0x01, []byte{0xaa}, "BCC"},
			{0x02, []byte{0xed}, "SEC"},
			{0x03, []byte{0x86}, "BCS"},
			{0x04, []byte{0x27}, "CLC"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-shift.bin",
		Description:     "Shift instructions",
		FinalState:      Registers{0xa0, 0xff, 0x00},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0x00}, "No error"},
			{0x01, []byte{0x1e}, "ASL"},
			{0x02, []byte{0x79}, "LSR"},
			{0x03, []byte{0x02, 0x05}, "ROL"},
			{0x05, []byte{0x40, 0xa0}, "ROR"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-stack.bin",
		Description:     "Stack instructions",
		FinalState:      Registers{0x93, 0xce, 0x93},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0x20}, "Set stack pointer"},
			{0x011f, []byte{0x8f, 0xec}, "Push accumulator to stack"},
			{0x01, []byte{0xa6, 0x0e}, "Pull accumulator from stack"},
			{0x03, []byte{0xce, 0x93}, "TXA and TYA"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-adc.bin",
		Description:     "ADC instruction",
		FinalState:      Registers{0xe3, 0x08, 0x96},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0x00}, "No error"},
			{0x01, []byte{0xaf}, "ADC Immediate"},
			{0x02, []byte{0x55}, "ADC Absolute"},
			{0x03, []byte{0x4b}, "ADC Absolute,X"},
			{0x04, []byte{0x80}, "ADC Absolute,Y"},
			{0x05, []byte{0xca}, "ADC Zero Page"},
			{0x06, []byte{0x2e}, "ADC Zero Page,X"},
			{0x07, []byte{0xad}, "ADC indexed indirect"},
			{0x08, []byte{0xe3}, "ADC indirect indexed"},
		},
	},

	Program{
		MachineCodeFile: "../../asm/test-overflow.bin",
		Description:     "Overflow flag",
		FinalState:      Registers{0xd4, 0x00, 0x00},
		FinalMemory: []MemoryMatch{
			{0x00, []byte{0x00}, "No error"},
			{0x01, []byte{0x24}, "BVS"},
			{0x02, []byte{0x8e}, "BVC"},
			{0x03, []byte{0xd4}, "CLV"},
		},
	},
}
