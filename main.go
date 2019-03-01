package main

import (
	"fmt"
)

var test1 = Program{
	MachineCodeFile: "asm/test1",
	FinalState:      Registers{0x01, 0x10, 0x22},
}

func main() {
	fmt.Println(test1.Check())
}
