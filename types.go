package main

type targetFunction func(MoveRequest) Coordinate

type Coordinate struct {
	X int
	Y int
}

type MoveResponse struct {
	Move  string
	Shout string
}

type EndRequest struct {
	Game  string
	Turn  string
	Board string
	You   string
}

// This is the request structure from the gameserver -
// We can pass the whole thing to any functions which are
// computing moves.
type MoveRequest struct {
	Game struct {
		ID string `json:"id"`
	} `json:"game"`
	Turn  int `json:"turn"`
	Board struct {
		Height int `json:"height"`
		Width  int `json:"width"`
		Food   []struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"food"`
		Snakes []struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Health int    `json:"health"`
			Body   []struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"body"`
			Shout string `json:"shout"`
		} `json:"snakes"`
	} `json:"board"`
	You struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Health int    `json:"health"`
		Body   []struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"body"`
		Shout string `json:"shout"`
	} `json:"you"`
}

type StartRequest struct {
	Game  string `json:"game,omitempty"`
	Turn  string `json:"turn,omitempty"`
	Board string `json:"board,omitempty"`
	You   string `json:"you,omitempty"`
}

type StartResponse struct {
	Color    string
	Headtype string
	Tailtype string
}
