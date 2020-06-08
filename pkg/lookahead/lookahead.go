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
	// for permutation:	(in parallel?)
	// - advance board(moves)
	// return {left: [board1, board2], right: []...} etc

}

func advanceBoard(state structs.Board, moves map[string]string) {
	
	// take a gamestate and a mapping of snakes:moves
	// {state, [snake1: left, snake2: right]
	// simulate the game for one tick
	// return the updated gamestate

}

func NextBoards(state structs.MoveRequest, depth int) <-chan string {
	c := make(chan string)

	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan string) {
		defer close(c) // Once the iteration function is finished, we close the channel

		MakeMoves(c, state, depth) // We start by feeding it the current state of the game and a depth
	}(c)

	return c // Return the channel to the calling function
}

// AddLetter adds a letter to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddLetter once again
// to add yet another letter to the new combination in case length allows it
// Return: `moves`
func MakeMoves(c chan string, state structs.MoveRequest, depth int) {
	if depth <= 0 {
		return
	}

	moves = simple_moves.PossibleMoves(state)
	if len(moves) == 0 {
		return
	}


	for _, move := range moves {
		newState = applyMovesToState(move, state)
		c <- newState.Board
		MakeMoves(c, newState, depth - 1)
	}

	// 4 ryan 
}

// moves maps the snake ID to the series of moves that it'll make in `state`
func applyMovesToState(moves map{str, []string}, state structs.MoveRequest) structs.MoveRequest {
	// This is confusing because state.Board contains all of the snakes, including you
	// But the only way you know which one is you is from state.You
	// So it really needs to be updated in two places. Like you need to apply the move to the board snakes and also You.
	// I'm hungry.

	// 4 sam
}