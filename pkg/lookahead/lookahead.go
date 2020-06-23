package lookahead

import (
	"math/rand"
	"time"

	"github.com/SamWheating/battlesnake2020/pkg/heuristics"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

func Lookahead(state structs.MoveRequest, depth int, heuristic heuristics.Heuristic) string {
	// // moves := NextBoards(state, depth)	// all the possible next boards: {left: [board1, board2...], right: []}
	// scores := make(map[string]int) // {left: 100, right: 80, down: 110}
	// for _, move := range moves {
	// 	for board := range moves[move] {
	// 		scores[move] += heuristic.score(state)
	// 	}
	// }

	return ""
}

// func NextBoards(state structs.MoveRequest, depth int){

// 	calculate all possible permutations
// 	for permutation:	(in parallel?)
// 	- advance board(moves)
// 	return {left: [board1, board2], right: []...} etc

// }

func advanceBoard(state structs.Board, moves map[string]string) {

	// take a gamestate and a mapping of snakes:moves
	// {state, [snake1: left, snake2: right]
	// simulate the game for one tick
	// return the updated gamestate

}

// FWIW: https://stackoverflow.com/a/19249957 inspired a lot of what I did here.

// NextBoards returns some boards.
func NextBoards(state structs.MoveRequest, depth int) <-chan structs.Board {
	c := make(chan structs.Board)

	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan structs.Board) {
		defer close(c) // Once the iteration function is finished, we close the channel

		makeMoves(c, state, depth) // We start by feeding it the current state of the game and a depth
	}(c)

	return c // Return the channel to the calling function
}

// This new combination is passed on to the channel before we call makeMoves once again
// Return: `moves` (will be passed to the channel and )
func makeMoves(c chan structs.Board, state structs.MoveRequest, depth int) {
	if depth <= 0 {
		return
	}

	moves := []string{"up", "down", "left", "right"}
	snakes := state.Board.Snakes
	snakeMoves := make([]map[string]string, 0)
	for _, snake := range snakes {
		var snakeMove = make(map[string]string)
		rand.Seed(time.Now().Unix())
		move := moves[rand.Intn(len(moves))]
		snakeMove[snake.ID] = move
		snakeMoves = append(snakeMoves, snakeMove)
	}

	newState := applyMovesToState(snakeMoves, state)
	c <- newState.Board
	makeMoves(c, newState, depth-1)

	// 4 ryan
}

// moves maps the snake ID to the series of moves that it'll make in `state`
func applyMovesToState(moves []map[string]string, state structs.MoveRequest) structs.MoveRequest {
	_ = moves
	state.Turn++
	return state
	// This is confusing because state.Board contains all of the snakes, including you
	// But the only way you know which one is you is from state.You
	// So it really needs to be updated in two places. Like you need to apply the move to the board snakes and also You.
	// I'm hungry.
	// MoveRequest turn += 1

	// 4 sam
}
