package lookahead

import (
	"github.com/SamWheating/battlesnake2020/pkg/heuristics"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

func Lookahead(state structs.MoveRequest, depth int, heuristics.Heuristic heuristic) string {
	moves := NextBoards()
	scores := make(map[string]int)
	my_id = state.You.ID
	for _, move := range moves {
		for board := range moves[move] {
			scores[move] += heuristic.score(board, my_id)
		}
	}
}
