// Define move-computing functions in here to heep main.go clean

package main

import (
	"fmt"
)

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

func followTail(state MoveRequest) string {

	safe_moves := make(map[string]bool)
	for _, move := range all_moves {
		if (isSafeMove(state, move)){
			safe_moves[move] = true
		} else {
			safe_moves[move] = false
		}
	}

	fmt.Println(safe_moves)

	current := state.You.Body[0]
	target := state.You.Body[len(state.You.Body) - 1]
	fmt.Println(target)

	if((target.X < current.X) && (safe_moves["left"])){
		return "left"
	} else if((target.Y > current.Y) && (safe_moves["down"])) {
		return "down"
	} else if((target.X > current.X) && (safe_moves["right"])) {
		return "right"
	} else if((target.Y < current.Y) && (safe_moves["up"])){
		return "up"
	} else {
		return "down"
	}
}

func isOutOfBounds(state MoveRequest, move string ) bool {
	current := state.You.Body[0]
	next_x := current.X + moves[move]["X"]
	next_y := current.Y + moves[move]["Y"]
	if (next_x >= state.Board.Width || next_x < 0) {
		return true
	}
	if (next_y >= state.Board.Height || next_y < 0) {
		return true
	}
	return false
}

func isSelfCollision(state MoveRequest, move string ) bool {
	current := state.You.Body
	next_x := current[0].X + moves[move]["X"]
	next_y := current[0].Y + moves[move]["Y"]
	for _, element := range current {
		if(element.X == next_x && element.Y == next_y){
			return true
		} 
	}
	return false
}

func isSafeMove(state MoveRequest, move string) bool {
	//if(!(isOutOfBounds(state, move)) && !(isSelfCollision(state, move))){
	if(!(isSelfCollision(state, move)) && !(isOutOfBounds(state, move))){
		return true
	}
	return false
}