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
	FinalMemory:     []MemoryMatch{MemoryMatch{0xa0, []byte{0x01, 0x10, 0x22}}},
}

var test4 = Program{
	MachineCodeFile: "asm/test4",
	Description:     "TAX/TAY",
	FinalState:      Registers{0x01, 0x01, 0x01},
}

var test8 = Program{
	MachineCodeFile: "asm/test8",
	Description:     "LD? Zero Page",
	FinalState:      Registers{0x22, 0x10, 0x10},
	FinalMemory:     []MemoryMatch{MemoryMatch{0x00, []byte{0x01, 0x10, 0x22}}},
}

var testLD = Program{
	MachineCodeFile: "asm/test-lda",
	Description:     "LD? All Addressing Modes",
	FinalState:      Registers{0xa0, 0xa2, 0x02},
	FinalMemory: []MemoryMatch{
		MemoryMatch{0x00, []byte{0x22, 0x01, 0x10, 0xa9, 0xa2, 0xa0, 0xa2, 0xa9, 0xa0, 0xa2}},
	},
}

func main() {
	fmt.Println(test1.Description, ":", test1.Check())
	fmt.Println(test2.Description, ":", test2.Check())
	fmt.Println(test4.Description, ":", test4.Check())
	fmt.Println(test8.Description, ":", test8.Check())
	fmt.Println(testLD.Description, ":", testLD.Check())
}
