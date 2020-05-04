package moves_test

import (
	"testing"
	"github.com/SamWheating/battlesnake2020/pkg/moves"
)

func TestAbs(t *testing.T){
	got := moves.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d, expected 1", got)
	}
	got = moves.Abs(10)
	if got != 10 {
		t.Errorf("Abs(-1) = %d, expected 10", got)
	}
}

