package main // import "github.com/jechols/turingo"

import (
	"fmt"

	"github.com/jechols/turingo/machine"
)

func main() {
	var m = machine.New()

	// Destroy All Software initial "flip symbols" example
	m.AddInstruction("s1", 'B', 'X', 'R', "s2")
	m.AddInstruction("s2", 'B', 'B', 'L', "s3")
	m.AddInstruction("s3", 'X', 'B', 'R', "s4")
	m.AddInstruction("s4", 'B', 'B', 'L', "s1")

	fmt.Println("DAS: Toggle B->X")
	fmt.Println("----------------")
	m.Run("s1", 8, func() {
		fmt.Println(m.String())
	})
	fmt.Println()

	// DAS adding machine
	m = machine.New()

	// Initial state: the problem of 2 + 3
	m.AddInstruction("s1", 'B', '(', 'R', "s2")
	m.AddInstruction("s2", 'B', '1', 'R', "s3")
	m.AddInstruction("s3", 'B', '1', 'R', "s4")
	m.AddInstruction("s4", 'B', '+', 'R', "s5")
	m.AddInstruction("s5", 'B', '1', 'R', "s6")
	m.AddInstruction("s6", 'B', '1', 'R', "s7")
	m.AddInstruction("s7", 'B', '1', 'R', "s8")
	m.AddInstruction("s8", 'B', ')', 'R', "s9")

	// Find the plus
	m.AddInstruction("s9", 'B', 'B', 'L', "s9")
	m.AddInstruction("s9", '1', '1', 'L', "s9")
	m.AddInstruction("s9", ')', ')', 'L', "s9")

	// Remove and combine
	m.AddInstruction("s9", '+', '1', 'R', "s10")
	m.AddInstruction("s10", '1', '1', 'R', "s10")
	m.AddInstruction("s10", ')', 'B', 'L', "s11")
	m.AddInstruction("s11", '1', ')', 'L', "done")

	fmt.Println("DAS: Add 2+3")
	fmt.Println("------------")
	m.Run("s1", -1, func() {
		fmt.Println(m.String())
	})
}
