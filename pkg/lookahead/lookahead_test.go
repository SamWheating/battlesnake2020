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
