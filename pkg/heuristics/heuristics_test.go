package heuristics_test

import (
	"github.com/SamWheating/battlesnake2020/pkg/heuristics"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
	"testing"
)

func TestHunger(t *testing.T) {
	snake := new(structs.Snake)
	snake.ID = "test_snake"
	snake.Health = 90
	board := new(structs.Board)
	board.Snakes = []structs.Snake{*snake}

	got := heuristics.Hunger(*board, "test_snake")
	if got != 90 {
		t.Errorf("Hunger was %d, expected 90", got)
	}
}
