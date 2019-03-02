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

var test3 = Program{
	MachineCodeFile: "asm/test3",
	Description:     "LD? Zero Page",
	FinalState:      Registers{0x22, 0x01, 0x10},
	FinalMemory:     []MemoryMatch{MemoryMatch{0xa0, []byte{0x01, 0x10, 0x22}}},
}

var test4 = Program{
	MachineCodeFile: "asm/test4",
	Description:     "TAX/TAY",
	FinalState:      Registers{0x01, 0x01, 0x01},
}

var test5 = Program{
	MachineCodeFile: "asm/test5",
	Description:     "LD? Absolute",
	FinalState:      Registers{0xa9, 0xa2, 0xa0},
}

func main() {
	fmt.Println(test1.Description, ":", test1.Check())
	fmt.Println(test2.Description, ":", test2.Check())
	fmt.Println(test3.Description, ":", test3.Check())
	fmt.Println(test4.Description, ":", test4.Check())
	fmt.Println(test5.Description, ":", test5.Check())
}
