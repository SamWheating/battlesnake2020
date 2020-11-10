package heuristics

import (
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

// Heuristic functions are applied to a board and return a Float value.

type Heuristic interface {
	score(structs.MoveRequest) float64
}

func Hunger(state structs.MoveRequest) int {
	return state.You.Health
}

// Calculates the total reachable spaces from our snake's head
// Heuristic function (assigns a score to a hypothetical game state)
func HeadRoom(board structs.Board, you string) int {
	head_x := -1
	//head_y := -1
	youSnake := structs.Snake{}
	for _, snake := range board.Snakes {
		if snake.ID == you {
			youSnake = snake
			head_x = snake.Body[0].X
			//head_y = snake.Body[0].Y
		}
	}
	// penalize the snake for dying in this turn
	if head_x == -1 {
		return -5 // TUNE THIS MAGIC CONSTANT
		panic("got here")
	}

	// Initialize a width x height array of false
	boolboard := make([][]bool, board.Width)
	for i := range boolboard {
		boolboard[i] = make([]bool, board.Height)
	}

	for _, snake := range board.Snakes {
		for i, coord := range snake.Body {
			if !(snake.ID == you && i == 0) {
				boolboard[coord.X][coord.Y] = true
			}
			if snake.ID != you {
				head := snake.Body[0]
				if IsInBounds(boolboard, head.Up()) {
					boolboard[head.Up().X][head.Up().Y] = true
				}
				if IsInBounds(boolboard, head.Left()) {
					boolboard[head.Left().X][head.Left().Y] = true
				}
				if IsInBounds(boolboard, head.Right()) {
					boolboard[head.Right().X][head.Right().Y] = true
				}
				if IsInBounds(boolboard, head.Down()) {
					boolboard[head.Down().X][head.Down().Y] = true
				}
			}
		}
	}

	//boolboard[head_x][head_y] = false
	score := FloodFill(boolboard, youSnake.Body[0])

	if youSnake.Health < 40 {
		score -= 10 * (40 - youSnake.Health)
	}

	return score
}

func FloodFill(boardState [][]bool, coord structs.Coordinate) int {

	// Stack-based recursive implementation (four-way)

	// One implicitly stack-based (recursive) flood-fill implementation (for a two-dimensional array) goes as follows:

	// Flood-fill (node, target-color, replacement-color):
	//  1. If square is occupied, return zero
	//  3. Else Set the square to occupied AND count += 1
	//  4. Count += Flood-fill (one step to the south of node, target-color, replacement-color).
	//     Count += Flood-fill (one step to the north of node, target-color, replacement-color).
	//     Count += Flood-fill (one step to the west of node, target-color, replacement-color).
	//     Count += Flood-fill (one step to the east of node, target-color, replacement-color).
	//  5. Return Count.

	// make a copy of boardstate so that the original can be reused

	count := 0

	// Out of the board is bad!
	if coord.X >= len(boardState) || coord.X < 0 {
		return 0
	}
	if coord.Y >= len(boardState[0]) || coord.Y < 0 {
		return 0
	}

	if boardState[coord.X][coord.Y] {
		return 0
	}

	boardState[coord.X][coord.Y] = true
	count += 1

	count += FloodFill(boardState, coord.Left())
	count += FloodFill(boardState, coord.Right())
	count += FloodFill(boardState, coord.Up())
	count += FloodFill(boardState, coord.Down())
	return count
}

func IsInBounds(board [][]bool, coord structs.Coordinate) bool {
	if coord.X < 0 || coord.Y < 0 {
		return false
	}
	if coord.X >= len(board) || coord.Y >= len(board[0]) {
		return false
	}
	return true
}
