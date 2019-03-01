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

func main() {
	fmt.Println(test1.Description, ":", test1.Check())
	fmt.Println(test2.Description, ":", test2.Check())
}
