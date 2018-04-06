package machine // import "github.com/jechols/turingo/machine"

import (
	"errors"
	"fmt"
)

// Machine is our turing machine
type Machine struct {
	state string
	tape  []rune
	head  int

	instructions map[condition]operation
}

// New produces a new turing machine for use
func New() *Machine {
	var m = &Machine{instructions: make(map[condition]operation)}
	m.growTape()
	m.head = 127
	return m
}

// A condition is essentially the "if" part of the Machine: if we are in a
// given state and see a given value at the current head, this condition is
// considered to be satisfied
type condition struct {
	state string
	tVal  rune
}

// An operation tells us the new value to write, the direction to move (left or right), and the new state to set
type operation struct {
	val   rune
	dir   rune
	state string
}

// dirmap defines what the direction runes actually mean.  It would be simpler
// to make conditions just store a raw int, but runes help visualize the
// machine's instructions.
var dirmap = map[rune]int{'L': -1, 'R': 1}

func (m *Machine) AddInstruction(state string, seeVal rune, newValue rune, headDirection rune, newState string) error {
	if headDirection != 'R' && headDirection != 'L' {
		return errors.New("invalid direction")
	}
	var c = condition{state, seeVal}
	if _, ok := m.instructions[c]; ok {
		return errors.New("repeating condition: " + state + "/" + string(seeVal))
	}

	m.instructions[c] = operation{newValue, headDirection, newState}
	return nil
}

func (m *Machine) Run(state string, n int, itercb func()) error {
	m.state = state
	for n > 0 {
		itercb()
		m.growTape()
		n--
		var c = condition{m.state, m.tape[m.head]}
		var op, ok = m.instructions[c]
		if !ok {
			return errors.New("machine state has no valid instructions for continuing")
		}

		m.tape[m.head] = op.val
		m.head += dirmap[op.dir]
		m.state = op.state
	}

	// We call the iterator one extra time to get the final state
	itercb()

	return nil
}

// growTape checks to see if head is too far left or right, and adds data to
// the tape (adjusting head if necessary).  To avoid recreating tape data too
// often, we provision 256 bytes at a time.
func (m *Machine) growTape() {
	if m.head < 0 {
		var newTape = make([]rune, 256)
		m.tape = append(newTape, m.tape...)
		m.head += 256
		m.fillTape()
	}
	if m.head >= len(m.tape) {
		var newTape = make([]rune, 256)
		m.tape = append(m.tape, newTape...)
		m.fillTape()
	}
}

// fillTape just turns empty tape cells into the letter 'B' to represent the blank state
func (m *Machine) fillTape() {
	for i, val := range m.tape {
		if val == 0 {
			m.tape[i] = 'B'
		}
	}
}

func (m *Machine) String() string {
	// We print out only the nearest bytes, but in a way that attempts to show
	// useful context as the head pointer moves
	var prefix = "..."
	var tapeStart = (m.head / 10) * 10
	if tapeStart == m.head {
		prefix = ""
	}

	var suffix = "..."
	var tapeEnd = tapeStart + 15
	if tapeEnd >= len(m.tape) {
		tapeEnd = len(m.tape) - 1
		suffix = ""
	}

	var out = fmt.Sprintf("%6s: ", m.state) + prefix

	var tapeSlice = m.tape[tapeStart:tapeEnd]
	for i, r := range tapeSlice {
		if m.head == i+tapeStart {
			out += "<" + string(r) + ">"
		} else {
			out += " " + string(r) + " "
		}
	}

	out += suffix
	return out
}
