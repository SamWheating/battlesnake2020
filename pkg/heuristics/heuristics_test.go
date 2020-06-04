package heuristics_test

import (
	"testing"

	"github.com/SamWheating/battlesnake2020/pkg/heuristics"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

func TestHunger(t *testing.T) {

	//makes a mock gamestate
	snake := new(structs.Snake)
	snake.ID = "test_snake"
	snake.Health = 90
	board := new(structs.Board)
	board.Snakes = []structs.Snake{*snake}
	request := new(structs.MoveRequest)
	request.Board = *board
	request.You = *snake

	got := heuristics.Hunger(*request)
	if got != 90 {
		t.Errorf("Hunger was %d, expected 90", got)
	}
}

func TestFloodFillSingleSquare(t *testing.T) {
	board := [][]bool{
		{true, true, true},
		{true, false, true},
		{true, true, true},
	}

	coord := new(structs.Coordinate)
	coord.X = 1
	coord.Y = 1

	got := heuristics.FloodFill(board, *coord)
	if got != 1 {
		t.Errorf("Floodfilled to %d, expected 1", got)
	}
}

func TestFloodFillAllEmpty(t *testing.T) {
	board := [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}

	coord := new(structs.Coordinate)
	coord.X = 1
	coord.Y = 1

	got := heuristics.FloodFill(board, *coord)
	if got != 9 {
		t.Errorf("Floodfilled to %d, expected 9", got)
	}
}

func TestFloodFillSomeEmpty(t *testing.T) {
	board := [][]bool{
		{false, true, false},
		{false, true, false},
		{false, false, false},
	}

	coord := new(structs.Coordinate)
	coord.X = 0
	coord.Y = 0

	got := heuristics.FloodFill(board, *coord)
	if got != 7 {
		t.Errorf("Floodfilled to %d, expected 7", got)
	}
}

func TestHeadRoom(t *testing.T) {

	request := structs.MoveRequest{
		You: structs.Snake{
			Body: []structs.Coordinate{
				structs.Coordinate{
					X: 0,
					Y: 1,
				},
				structs.Coordinate{
					X: 1,
					Y: 1,
				},
			},
		},
		Board: structs.Board{
			Height: 2,
			Width:  2,
			Snakes: []structs.Snake{
				structs.Snake{
					Body: []structs.Coordinate{
						structs.Coordinate{
							X: 0,
							Y: 1,
						},
						structs.Coordinate{
							X: 1,
							Y: 1,
						},
					},
				},
			},
		},
	}

	got := heuristics.HeadRoom(request)
	if got != 3 {
		t.Errorf("Floodfilled to %d, expected 3", got)
	}
}

func TestHeadRoomFullBoard(t *testing.T) {

	request := structs.MoveRequest{
		You: structs.Snake{
			Body: []structs.Coordinate{
				structs.Coordinate{
					X: 0,
					Y: 1,
				},
				structs.Coordinate{
					X: 1,
					Y: 1,
				},
			},
		},
		Board: structs.Board{
			Height: 2,
			Width:  2,
			Snakes: []structs.Snake{
				structs.Snake{
					Body: []structs.Coordinate{
						structs.Coordinate{
							X: 0,
							Y: 1,
						},
						structs.Coordinate{
							X: 1,
							Y: 1,
						},
					},
				},
				structs.Snake{
					Body: []structs.Coordinate{
						structs.Coordinate{
							X: 0,
							Y: 0,
						},
						structs.Coordinate{
							X: 1,
							Y: 0,
						},
					},
				},
			},
		},
	}

	got := heuristics.HeadRoom(request)
	if got != 1 {
		t.Errorf("Floodfilled to %d, expected 3", got)
	}
}
