package machine

import (
	"strings"
	"testing"
)

func TestGrowTape(t *testing.T) {
	// Creating a new machine should start us in the middle of a tape
	var m = New()
	if string(m.tape) != strings.Repeat("B", 256) || m.head != 127 {
		t.Errorf("Expected an initialized machine, got a tape of %q and head of %d", m.tape, m.head)
	}

	m.head = -5
	m.growTape()
	if len(m.tape) != 512 && string(m.tape) != strings.Repeat("B", 512) {
		t.Errorf(`After growing a tape with a head at -5, we should see 512 blank bytes, but got %q`, string(m.tape))
	}
	if m.head != 251 {
		t.Errorf(`After growing an empty tape with a head at -5, new head should be 251, but it was %d`, m.head)
	}
}
