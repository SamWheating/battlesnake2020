package lookahead

import (
	"fmt"
	"github.com/SamWheating/battlesnake2020/pkg/heuristics"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
	"github.com/getlantern/deepcopy"
	"math/rand"
	"time"
)

func Lookahead(state structs.MoveRequest, depth int, count int) string {
	directions := []string{"left", "right", "up", "down"}
	scores := make(map[string][]int)
	for _, direction := range directions {
		scores[direction] = []int{}
	}
	moves := SampleRandomSnakeMoves(state.Board, depth, count)
	for _, move := range moves {
		board := ApplyMovesToBoard(move, state.Board)
		score := heuristics.HeadRoom(board, state.You.ID)
		direction := move[state.You.ID][0]
		scores[direction] = append(scores[direction], score)
	}

	fmt.Println("\n")
	max := -500.0
	choice := "left"
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
		fmt.Println(dir, dirScore)
	}
	fmt.Printf("go %s\n", choice)
	return choice
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
			moves := make([]string, depth)
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
func ApplyMovesToBoard(moves map[string][]string, board structs.Board) structs.Board {
	// snake1: [left, right, down]
	boardBoard := new(structs.Board)
	err := deepcopy.Copy(boardBoard, &board)
	if err != nil {
		fmt.Println(err)
	}
	newBoard := *boardBoard

	for i := range moves[newBoard.Snakes[0].ID] { // [left, right, down]
		snakes := []structs.Snake{}
		for j, snake := range newBoard.Snakes {
			next := snake.Body[0].Move(moves[snake.ID][i])
			newBoard.Snakes[j].Body = append([]structs.Coordinate{next}, snake.Body...)
			newBoard.Snakes[j].Health = snake.Health - 1
			if !CoordInList(snake.Body[0], newBoard.Food) {
				newBoard.Snakes[j].Body = newBoard.Snakes[j].Body[:len(newBoard.Snakes[j].Body)-1]
			} else {
				newBoard.Snakes[j].Health = 100
			}
			// only keep snakes which haven't starved or gone out of bounds
			if !IsOutOfBounds(newBoard, next) && !IsStarved(newBoard.Snakes[j]) && !HitOtherSnake(newBoard, next) {
				snakes = append(snakes, newBoard.Snakes[j])
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
	if snake.Health <= 10 { // Todo: this is a cheap hack for avoiding complete starvation
		return true
	}
	return false
}
