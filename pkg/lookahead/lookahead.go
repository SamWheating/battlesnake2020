package lookahead

import (
	"github.com/SamWheating/battlesnake2020/pkg/heuristics"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

func Lookahead(state structs.MoveRequest, depth int, heuristics.Heuristic heuristic) string {
	moves := NextBoards(state, depth)	// all the possible next boards: {left: [board1, board2...], right: []}
	scores := make(map[string]int)		// {left: 100, right: 80, down: 110}
	for _, move := range moves {
		for board := range moves[move] {
			scores[move] += heuristic.score(state)
		}
	}
}


func NextBoards(state structs.MoveRequest, depth int){

	// calculate all possible permutations
	// for permutation:
	// - advance board(moves)
	// return {left: [baord1, board2], right: []...} etc

}

func advanceBoard(state structs.Board, moves map[string]string) {
	
	// take a gamestate and a mapping of snakes:moves
	// {state, [snake1: left, snake2: right]
	// simulate the game for one tick
	// return the updated gamestate

}