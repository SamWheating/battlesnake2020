// Define move-computing functions in here to heep main.go clean

package main

import (
	"fmt"
)

func Abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func playItSafe(state MoveRequest) string {
	if state.You.Health > 30 {
		return shy(state)
	} else {
		return greedy(state)
	}
}

func moveToTarget(state MoveRequest, targetFunc targetFunction) string {
	safe_moves := make(map[string]bool)
	for _, move := range all_moves {
		if isSafeMove(state, move) {
			safe_moves[move] = true
		} else {
			safe_moves[move] = false
		}
	}

	current := state.You.Body[0]
	target := targetFunc(state)

	if (target.X < current.X) && (safe_moves["left"]) {
		return "left"
	} else if (target.Y > current.Y) && (safe_moves["down"]) {
		return "down"
	} else if (target.X > current.X) && (safe_moves["right"]) {
		return "right"
	} else if (target.Y < current.Y) && (safe_moves["up"]) {
		return "up"
	} else {
		for k, v := range safe_moves {
			if v {
				return k
			}
		}
		fmt.Println("no safe?")
		return "down"
	}
}

func closestFood(state MoveRequest) Coordinate {
	head := state.You.Body[0]
	all_food := state.Board.Food
	min := 100000
	var closest Coordinate
	for _, food := range all_food {
		distance := Abs(food.X-head.X) + Abs(food.Y-head.Y)
		if distance < min {
			min = distance
			closest.X = food.X
			closest.Y = food.Y
		}
	}
	return closest
}

var all_moves = [4]string{"left", "right", "up", "down"}

var moves = map[string]map[string]int{
	"up": map[string]int{
		"X": 0, "Y": -1},
	"down": map[string]int{
		"X": 0, "Y": 1},
	"left": map[string]int{
		"X": -1, "Y": 0},
	"right": map[string]int{
		"X": 1, "Y": 0},
}

// Uses a heuristic function to find the safest square on the board
func shy(state MoveRequest) string {
	return moveToTarget(state, safestSquare)
}

// Follows its own tail
func followTail(state MoveRequest) string {
	return moveToTarget(state, findTail)
}

// Always goes to the nearest food
func greedy(state MoveRequest) string {
	return moveToTarget(state, closestFood)
}

func possibleMoves(state MoveRequest) []Coordinate {
	current := state.You.Body[0]
	possible := []Coordinate{}
	for _, move := range all_moves {
		if isSafeMove(state, move) {
			var coord Coordinate
			coord.X = current.X + moves[move]["X"]
			coord.Y = current.Y + moves[move]["Y"]
			possible = append(possible, coord)
		}
	}
	return possible
}

func findTail(state MoveRequest) Coordinate {
	var tail Coordinate
	tail.X = state.You.Body[len(state.You.Body)-1].X
	tail.Y = state.You.Body[len(state.You.Body)-1].Y
	return tail
}

func safestSquare(state MoveRequest) Coordinate {
	best := 0
	var best_square Coordinate
	for _, move := range possibleMoves(state) {
		current := 0
		for _, snake := range state.Board.Snakes {
			for _, element := range snake.Body[:len(snake.Body)-1] {
				current += Abs(move.X-element.X) + Abs(move.Y-element.Y)
			}
		}
		if current > best {
			best = current
			best_square.X = move.X
			best_square.Y = move.Y
		}
	}
	return best_square
}

func isOutOfBounds(state MoveRequest, move string) bool {
	current := state.You.Body[0]
	next_x := current.X + moves[move]["X"]
	next_y := current.Y + moves[move]["Y"]
	if next_x >= state.Board.Width || next_x < 0 {
		return true
	}
	if next_y >= state.Board.Height || next_y < 0 {
		return true
	}
	return false
}

func isOtherCollision(state MoveRequest, move string) bool {
	snakes := state.Board.Snakes
	current := state.You.Body
	next_x := current[0].X + moves[move]["X"]
	next_y := current[0].Y + moves[move]["Y"]
	for _, snake := range snakes {
		for _, element := range snake.Body[:len(snake.Body)-1] {
			if element.X == next_x && element.Y == next_y {
				return true
			}
		}
	}
	return false
}

func isSafeMove(state MoveRequest, move string) bool {
	if !(isOtherCollision(state, move)) && !(isOutOfBounds(state, move)) {
		return true
	}
	return false
}
