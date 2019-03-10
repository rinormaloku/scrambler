package scrambler

import (
	"testing"
)

func TestScramble(t *testing.T) {
	var initial = "This is the initial text"
	var final = Scramble(initial)
	if initial == final {
		t.Errorf("Text is not scrambled!")
	}
}
