package main // import "github.com/jechols/turingo"

import (
	"fmt"

	"github.com/jechols/turingo/machine"
)

func main() {
	var m = machine.New()

	// Destroy All Software initial "flip symbols" example
	m.AddInstruction("s1", machine.Empty, '1', 'R', "s2")
	m.AddInstruction("s2", machine.Empty, machine.Empty, 'L', "s3")
	m.AddInstruction("s3", '1', machine.Empty, 'R', "s4")
	m.AddInstruction("s4", machine.Empty, machine.Empty, 'L', "s1")

	fmt.Println("DAS: Toggle 0->1")
	fmt.Println("----------------")
	m.Run("s1", 8, func() {
		fmt.Println(m.String())
	})
	fmt.Println()

	// DAS adding machine
	m = machine.New()

	// Initial state: the problem of 2 + 3
	m.AddInstruction("s1", machine.Empty, '(', 'L', "s2")
	m.AddInstruction("s2", machine.Empty, '1', 'L', "s3")
	m.AddInstruction("s3", machine.Empty, '1', 'L', "s4")
	m.AddInstruction("s4", machine.Empty, '+', 'L', "s5")
	m.AddInstruction("s5", machine.Empty, '1', 'L', "s6")
	m.AddInstruction("s6", machine.Empty, '1', 'L', "s7")
	m.AddInstruction("s7", machine.Empty, '1', 'L', "s8")
	m.AddInstruction("s8", machine.Empty, ')', 'L', "s9")

	// Find the plus
	m.AddInstruction("s9", machine.Empty, machine.Empty, 'R', "s9")
	m.AddInstruction("s9", '1', '1', 'R', "s9")
	m.AddInstruction("s9", ')', ')', 'R', "s9")

	// Remove and combine
	m.AddInstruction("s9", '+', '1', 'L', "s10")
	m.AddInstruction("s10", '1', '1', 'L', "s10")
	m.AddInstruction("s10", ')', machine.Empty, 'R', "s11")
	m.AddInstruction("s11", '1', ')', 'R', "done")

	fmt.Println("DAS: Add 2+3")
	fmt.Println("------------")
	m.Run("s1", -1, func() {
		fmt.Println(m.String())
	})
}
