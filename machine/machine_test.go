package machine

import (
	"strings"
	"testing"
)

func TestGrowTape(t *testing.T) {
	// Creating a new machine should start us in the middle of a tape
	var m = New()
	if m.head != tapeGrow {
		t.Errorf("Expected an initialized machine to have a head at %d, but it was %d", tapeGrow, m.head)
	}
	if len(m.tape) != tapeGrow * 2 {
		t.Errorf("Expected an initialized machine's tape to be %d bytes, but it was %d", tapeGrow*2, len(m.tape))
	}
	if string(m.tape) != strings.Repeat("0", tapeGrow * 2) {
		t.Errorf("Expected an initialized machine to be filled with zeroes")
	}

	m.head = -5
	m.growTape()
	if len(m.tape) != tapeGrow*3 && string(m.tape) != strings.Repeat("0", tapeGrow*3) {
		t.Errorf(`We should see tapeGrow*3 blank bytes, but got %d bytes: %q`, len(m.tape), string(m.tape))
	}
	if m.head != 251 {
		t.Errorf(`New head should be 251, but it was %d`, m.head)
	}
}
