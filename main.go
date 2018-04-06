package main // import "github.com/jechols/turingo"

import (
	"fmt"

	"github.com/jechols/turingo/machine"
)

func main() {
	var m = machine.New()

	m.AddInstruction("s1", 'B', 'X', 'R', "s2")
	m.AddInstruction("s2", 'B', 'B', 'L', "s3")
	m.AddInstruction("s3", 'X', 'B', 'R', "s4")
	m.AddInstruction("s4", 'B', 'B', 'L', "s1")
	m.Run("s1", 8, func() {
		fmt.Println(m.String())
	})
}
