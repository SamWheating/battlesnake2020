package simple_moves_test

import (
	"github.com/SamWheating/battlesnake2020/pkg/simple_moves"
	"testing"
)

func TestAbs(t *testing.T) {
	got := simple_moves.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d, expected 1", got)
	}
	got = simple_moves.Abs(10)
	if got != 10 {
		t.Errorf("Abs(-1) = %d, expected 10", got)
	}
}
