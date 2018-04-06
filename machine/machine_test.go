package machine

import (
	"strings"
	"testing"
)

func TestGrowTape(t *testing.T) {
	// Creating a new machine should start us in the middle of a tape
	var m = New()
	if string(m.tape) != strings.Repeat("B", tapeGrow) || m.head != 0 {
		t.Errorf("Expected an initialized machine, got a tape of %q and head of %d", m.tape, m.head)
	}

	m.head = -5
	m.growTape()
	if len(m.tape) != tapeGrow*2 && string(m.tape) != strings.Repeat("B", tapeGrow*2) {
		t.Errorf(`We should see tapeGrow*2 blank bytes, but got %d bytes: %q`, len(m.tape), string(m.tape))
	}
	if m.head != 251 {
		t.Errorf(`New head should be 251, but it was %d`, m.head)
	}
}
