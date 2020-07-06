package lookahead_test

import (
	"testing"

	"github.com/SamWheating/battlesnake2020/pkg/lookahead"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

func TestNextBoards(t *testing.T) {
	state := structs.MoveRequest{
		Game: structs.Game{
			ID: "666",
		},
		Turn: 4,
		You: structs.Snake{
			Body: []structs.Coordinate{
				{
					X: 0,
					Y: 1,
				}, {
					X: 1,
					Y: 1,
				},
			},
		},
		Board: structs.Board{
			Height: 2,
			Width:  2,
			Snakes: []structs.Snake{
				{
					Body: []structs.Coordinate{
						{
							X: 0,
							Y: 1,
						}, {
							X: 1,
							Y: 1,
						},
					},
				}, {
					Body: []structs.Coordinate{
						{
							X: 3,
							Y: 4,
						}, {
							X: 3,
							Y: 5,
						},
					},
				},
			},
		},
	}

	for board := range lookahead.NextBoards(state, 4) {
		t.Logf("The turn is...%+v", board)
	}
}

func TestIsStarved(t *testing.T) {
	snake := structs.Snake{
		Health: 100,
	}
	got := lookahead.IsStarved(snake)
	if got {
		t.Errorf("Says snake is starved, expected is not starved")
	}

	snake = structs.Snake{
		Health: 0,
	}
	got = lookahead.IsStarved(snake)
	if !got {
		t.Errorf("Says snake isn't starved, expected is starved")
	}
}

func TestIsOutOfBounds(t *testing.T) {
	board := structs.Board{
		Width:  5,
		Height: 5,
	}
	inBoundsCoords := []structs.Coordinate{
		structs.Coordinate{X: 0, Y: 0},
		structs.Coordinate{X: 4, Y: 4},
		structs.Coordinate{X: 2, Y: 2},
		structs.Coordinate{X: 1, Y: 3},
	}
	for _, coord := range inBoundsCoords {
		if lookahead.IsOutOfBounds(board, coord) {
			t.Errorf("Got [%d, %d] is out of bounds, expected is in bounds", coord.X, coord.Y)
		}
	}
	outOfBoundsCoords := []structs.Coordinate{
		structs.Coordinate{X: -1, Y: -1},
		structs.Coordinate{X: 5, Y: 0},
		structs.Coordinate{X: 5, Y: 5},
		structs.Coordinate{X: -2, Y: 3},
	}
	for _, coord := range outOfBoundsCoords {
		if !lookahead.IsOutOfBounds(board, coord) {
			t.Errorf("Got [%d, %d] is in bounds, expected out of bounds", coord.X, coord.Y)
		}
	}
}

func TestApplyMovesToBoardMovesSnake(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID:     "guy",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 1, Y: 1},
				},
			},
		},
		Food:   []structs.Coordinate{},
		Width:  5,
		Height: 5,
	}
	moves := map[string][]string{"guy": []string{"left"}}
	nextboard := lookahead.ApplyMovesToBoard(moves, board)
	if nextboard.Snakes[0].Body[0] != (structs.Coordinate{X: 0, Y: 1}) {
		t.Errorf("Snake Moved to [%d, %d], expected [0,1]", nextboard.Snakes[0].Body[0].X, nextboard.Snakes[0].Body[0].Y)
	}
}

func TestApplyMovesToBoardMovesSnakeStarves(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID:     "hungry",
				Health: 1,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 1, Y: 1},
				},
			},
		},
		Food:   []structs.Coordinate{},
		Width:  5,
		Height: 5,
	}
	moves := map[string][]string{"hungry": []string{"left"}}
	nextboard := lookahead.ApplyMovesToBoard(moves, board)
	if len(nextboard.Snakes) != 0 {
		t.Errorf("Snake survived, should have starved + been removed")
	}
}

func TestApplyMovesToBoardMovesSnakeHitsWall(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID:     "lost",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 0, Y: 0},
				},
			},
		},
		Food:   []structs.Coordinate{},
		Width:  5,
		Height: 5,
	}
	moves := map[string][]string{"lost": []string{"left"}}
	nextboard := lookahead.ApplyMovesToBoard(moves, board)
	if len(nextboard.Snakes) != 0 {
		t.Errorf("Snake survived, should have hit wall + been removed")
	}
}

func TestApplyMovesToBoardMovesSnakeHitsWallThenMoves(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID:     "lost",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 0, Y: 0},
				},
			},
			structs.Snake{
				ID:     "healthy",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 3, Y: 3},
				},
			},
		},
		Food:   []structs.Coordinate{},
		Width:  5,
		Height: 5,
	}
	moves := map[string][]string{
		"lost":    []string{"left", "left"},
		"healthy": []string{"left", "left"},
	}
	nextboard := lookahead.ApplyMovesToBoard(moves, board)
	if len(nextboard.Snakes) != 1 {
		t.Errorf("Snake survived, should have hit wall + been removed")
	}
}

func TestApplyMultipleMovesToBoard(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID:     "lost",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 0, Y: 0},
				},
			},
		},
		Food:   []structs.Coordinate{},
		Width:  5,
		Height: 5,
	}
	moves := map[string][]string{"lost": []string{"right", "right"}}
	nextboard := lookahead.ApplyMovesToBoard(moves, board)
	if nextboard.Snakes[0].Body[0] != (structs.Coordinate{X: 2, Y: 0}) {
		t.Errorf("Snake Moved to [%d, %d], expected [2,0]", nextboard.Snakes[0].Body[0].X, nextboard.Snakes[0].Body[0].Y)
	}
}

func TestApplyMultipleMovesToMultipleSnakes(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID:     "snake1",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 0, Y: 0},
				},
			},
			structs.Snake{
				ID:     "snake2",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 3, Y: 3},
				},
			},
		},
		Food:   []structs.Coordinate{},
		Width:  5,
		Height: 5,
	}
	moves := map[string][]string{
		"snake1": []string{"right", "right"},
		"snake2": []string{"left", "left"},
	}
	nextboard := lookahead.ApplyMovesToBoard(moves, board)
	if nextboard.Snakes[0].Body[0] != (structs.Coordinate{X: 2, Y: 0}) {
		t.Errorf("Snake Moved to [%d, %d], expected [2,0]", nextboard.Snakes[0].Body[0].X, nextboard.Snakes[0].Body[0].Y)
	}
	if nextboard.Snakes[1].Body[0] != (structs.Coordinate{X: 1, Y: 3}) {
		t.Errorf("Snake Moved to [%d, %d], expected [1,3]", nextboard.Snakes[1].Body[0].X, nextboard.Snakes[1].Body[0].Y)
	}
}

func TestApplyMovesToBoardMovesOneSnakeHitsWall(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID:     "lost",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 0, Y: 0},
				},
			},
			structs.Snake{
				ID:     "guy",
				Health: 100,
				Body: []structs.Coordinate{
					structs.Coordinate{X: 1, Y: 1},
				},
			},
		},
		Food:   []structs.Coordinate{},
		Width:  5,
		Height: 5,
	}
	moves := map[string][]string{"lost": []string{"left"}, "guy": []string{"right"}}
	nextboard := lookahead.ApplyMovesToBoard(moves, board)
	if len(nextboard.Snakes) != 1 {
		t.Errorf("%d Snakes survived, expected 1", len(nextboard.Snakes))
	}
	if nextboard.Snakes[0].ID == "lost" {
		t.Errorf("The wrong snake died")
	}
}

func TestGetSnakeMoves(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID: "snake1",
			},
		},
	}
	moves := lookahead.GetSnakeMoves(board, 1)
	if len(moves) != 4 {
		t.Errorf("Actual output not as expected")
	}
}

// func TestGetSnakeMovesDepthTwo(t *testing.T) {
// 	board := structs.Board{
// 		Snakes: []structs.Snake{
// 			structs.Snake{
// 				ID: "snake1",
// 			},
// 		},
// 	}
// 	moves := lookahead.GetSnakeMoves(board, 2)
// 	if len(moves) != 16 {
// 		t.Errorf("Should have generated 16 scenarios, generated %d", len(moves))
// 	}
// }

func TestSampleRandomSnakeMoves(t *testing.T) {
	board := structs.Board{
		Snakes: []structs.Snake{
			structs.Snake{
				ID: "snake1",
			},
			structs.Snake{
				ID: "snake2",
			},
		},
	}
	moves := lookahead.SampleRandomSnakeMoves(board, 4, 4)
	if len(moves) != 4 {
		t.Errorf("Should have generated 4 scenarios, generated %d", len(moves))
	}
	if len(moves[0]["snake1"]) != 4 {
		t.Errorf("Should have generated 4 moves per snake-scenario, generated %d", len(moves))
	}
}
