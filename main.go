package main // import "github.com/jechols/turingo"

import (
	"fmt"
	"strings"

	"github.com/jechols/turingo/machine"
)

func main() {
	var m = machine.New()

	// Destroy All Software initial "flip symbols" example
	m.AddInstruction("s1", machine.Empty, '1', 'R', "s2")
	m.AddInstruction("s2", machine.Empty, machine.Empty, 'L', "s3")
	m.AddInstruction("s3", '1', machine.Empty, 'R', "s4")
	m.AddInstruction("s4", machine.Empty, machine.Empty, 'L', "s1")
	printMachine(m, "DAS: Toggle 0->1", "s1", 8)

	// Flip symbols with noops
	m = machine.New()
	m.AddInstruction("s1", machine.Empty, '1', machine.NoOp, "s2")
	m.AddInstruction("s2", '1', machine.Empty, machine.NoOp, "s1")
	printMachine(m, "DAS: Toggle 0->1 with noops", "s1", 8)

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
	m.AddInstruction("s11", '1', ')', 'R', machine.StateComplete)

	printMachine(m, "DAS: Add 2+3", "s1", -1)

	// Three-state busy beaver
	m = machine.New()
	m.AddInstruction("A", '0', '1', 'R', "B")
	m.AddInstruction("A", '1', '1', 'L', "C")
	m.AddInstruction("B", '0', '1', 'L', "A")
	m.AddInstruction("B", '1', '1', 'R', "B")
	m.AddInstruction("C", '0', '1', 'L', "B")
	m.AddInstruction("C", '1', '1', machine.NoOp, machine.StateComplete)
	printMachine(m, "Wikipedia: Three-state busy beaver", "A", -1)
}

func printMachine(m *machine.Machine, title string, startingState string, iterations int) {
	fmt.Println()
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	var err = m.Run(startingState, iterations, func() {
		fmt.Println(m.String())
	})
	fmt.Println()
	if err != nil {
		fmt.Println("Error running machine:", err)
	}
}
