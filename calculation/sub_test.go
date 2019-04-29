package calculation

import "testing"

func TestSub(t *testing.T) {
	total := Sub(5, 5)
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
	}
}
