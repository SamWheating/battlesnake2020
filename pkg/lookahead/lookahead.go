package lookahead

import (
	//"math/rand"
	//"time"
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

	// 4 ryan

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

	// if depth <= 0 {
	// 	return
	// }

	// moves := []string{"up", "down", "left", "right"}
	// snakes := state.Board.Snakes
	// snakeMoves := make(map[string]string)
	// for _, snake := range snakes {
	// 	rand.Seed(time.Now().Unix())
	// 	move := moves[rand.Intn(len(moves))]
	// 	snakeMoves[snake.ID] = move
	// 	snakeMoves = append(snakeMoves, snakeMove)
	// }

	// newBoard := applyMovesToBoard(snakeMoves, state)
	// c <- newBoard
	// makeMoves(c, newBoard, depth-1)

	// 4 ryan
}

// getSnakeMoves generates a tree of all possible combinations of moves for each snake
// depth specifies the number of moves in the future.
// I think this is referred to as a triple cartesian product
// return will look like:
// [
//	{snake1: ["left", "left"], "snake2": ["left", "left"]},
//  {snake1: ["left", "down"], "snake2": ["left", "left"]}
// ]
// With all possible combinations of moves. The size of return will be 4^<num snakes>^<depth>
func GetSnakeMoves(board structs.Board, depth int) []map[string][]string {
	snakemoves := []map[string][]string{}
	directions := []string{"up", "down", "left", "right"}
	for i := 0; i < depth; i++ {
		for _, snake := range board.Snakes {
			for _, direction := range directions {
				snakemoves = append(snakemoves, map[string][]string{snake.ID: []string{direction}})
			}
		}
	}
	return snakemoves
}

// applyMovesToBoard applies a set of moves to a board, thus advancing the state of the game by one tick.
// moves maps the snake ID to the series of moves that it'll make in `state`.
// Note: the spawning of food is not accounted for here.
// Operations are carried out in the following order:
//   1) Advance position
//   2) Subtract hunger
//   3) Eat food
//   4) Check for wall collisions + starvations
//   5) Check for snake-on-snake collisions (TODO)
func ApplyMovesToBoard(moves map[string]string, board structs.Board) structs.Board {

	snakes := []structs.Snake{}

	for _, snake := range board.Snakes {
		next := snake.Body[0].Move(moves[snake.ID])
		snake.Body = append([]structs.Coordinate{next}, snake.Body...)
		snake.Health = snake.Health - 1
		if !CoordInList(snake.Body[0], board.Food) {
			snake.Body = snake.Body[:len(snake.Body)-1]
		} else {
			snake.Health = 100
		}
		// only keep snakes which haven't starved or gone out of bounds
		if !IsOutOfBounds(board, next) && !IsStarved(snake) {
			snakes = append(snakes, snake)
		}
	}
	// update snakes on the board to exclude dead snakes
	board.Snakes = snakes
	return board
}

func IsOutOfBounds(board structs.Board, head structs.Coordinate) bool {
	if head.X >= board.Width || head.X < 0 {
		return true
	}
	if head.Y >= board.Height || head.Y < 0 {
		return true
	}
	return false
}

func IsStarved(snake structs.Snake) bool {
	if snake.Health <= 0 {
		return true
	}
	return false
}
