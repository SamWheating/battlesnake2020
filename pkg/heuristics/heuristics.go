package heuristics

import (
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

// Heuristic functions are applied to a board and return a Float value.

type Heuristic interface {
	score(structs.Board, string) float64
}

func Hunger(b structs.Board, id string) int {
	for _, snake := range b.Snakes {
		if snake.ID == id {
			return snake.Health
		}
	}
	return 0
}
