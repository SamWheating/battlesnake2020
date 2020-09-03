package lookahead

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SamWheating/battlesnake2020/pkg/heuristics"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

func Lookahead(state structs.MoveRequest, depth int) string {
	directions := []string{"left", "right", "up", "down"}
	scores := make(map[string][]int)
	for _, direction := range directions {
		scores[direction] = []int{}
	}
	count := 20 // TUNE THIS MAGIC CONSTANT
	moves := SampleRandomSnakeMoves(state.Board, depth, count)
	for _, move := range moves {
		board := ApplyMovesToBoard(move, state.Board)
		score := heuristics.HeadRoom(board, state.You.ID)
		direction := move[state.You.ID][0]
		scores[direction] = append(scores[direction], score)
	}

	max := 0.0
	choice := "left"
	fmt.Println(moves)
	fmt.Println(scores)
	for dir, all_scores := range scores {
		total := 0
		for _, score := range all_scores {
			total += score
		}
		dirScore := float64(total) / float64(len(all_scores))
		if dirScore > max {
			choice = dir
			max = dirScore
		}
	}
	return choice
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

// SampleRandomSnakeMoves generates <count> possible combinations of <depth> moves for each snake
func SampleRandomSnakeMoves(board structs.Board, depth int, count int) []map[string][]string {
	rand.Seed(time.Now().Unix())
	snakemoves := make([]map[string][]string, count)
	directions := []string{"up", "down", "left", "right"}
	for i := 0; i < count; i++ {
		scenario := map[string][]string{}
		for _, snake := range board.Snakes {
			moves := make([]string, 4)
			for j := 0; j < depth; j++ {
				moves[j] = directions[rand.Intn(len(directions))]
			}
			scenario[snake.ID] = moves
		}
		snakemoves[i] = scenario
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
func ApplyMovesToBoard(moves map[string][]string, board structs.Board) structs.Board {
	// snake1: [left, right, down]
	// TODO!!!! Make a local copy of the board so we don't corrupt the original object

	newBoard := *new(structs.Board)
	newBoard.Height = board.Height
	newBoard.Width = board.Width
	newBoard.Snakes = []structs.Snake{}
	newBoard.Food = []structs.Coordinate{}

	for i := range moves[board.Snakes[0].ID] { // [left, right, down]
		snakes := []structs.Snake{}
		for j, snake := range board.Snakes {
			next := snake.Body[0].Move(moves[snake.ID][i])
			board.Snakes[j].Body = append([]structs.Coordinate{next}, snake.Body...)
			board.Snakes[j].Health = snake.Health - 1
			if !CoordInList(snake.Body[0], board.Food) {
				board.Snakes[j].Body = board.Snakes[j].Body[:len(board.Snakes[j].Body)-1]
			} else {
				board.Snakes[j].Health = 100
			}
			// only keep snakes which haven't starved or gone out of bounds
			if !IsOutOfBounds(board, next) && !IsStarved(board.Snakes[j]) && !HitOtherSnake(board, next) {
				snakes = append(snakes, board.Snakes[j])
			}
		}
		newBoard.Snakes = snakes
	}
	// update snakes on the board to exclude dead snakes
	return newBoard
}

// TODO: include logic of snake on snake collisions w.r.t size
func HitOtherSnake(board structs.Board, head structs.Coordinate) bool {
	count := 0
	for _, snake := range board.Snakes {
		for _, coord := range snake.Body {
			if coord.X == head.X && coord.Y == head.Y {
				count += 1
			}
		}
	}
	return count > 1
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
