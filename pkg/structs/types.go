package structs

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Left goes left
func (c Coordinate) Left() Coordinate {
	result := Coordinate{
		X: c.X - 1,
		Y: c.Y,
	}
	return result
}

func (c Coordinate) Right() Coordinate {
	result := Coordinate{
		X: c.X + 1,
		Y: c.Y,
	}
	return result
}

//// prod mode
// func (c Coordinate) Up() Coordinate {
// 	result := Coordinate{
// 		X: c.X,
// 		Y: c.Y + 1,
// 	}
// 	return result
// }

// func (c Coordinate) Down() Coordinate {
// 	result := Coordinate{
// 		X: c.X,
// 		Y: c.Y - 1,
// 	}
// 	return result
// }

// dev mode
func (c Coordinate) Up() Coordinate {
	result := Coordinate{
		X: c.X,
		Y: c.Y - 1,
	}
	return result
}

func (c Coordinate) Down() Coordinate {
	result := Coordinate{
		X: c.X,
		Y: c.Y + 1,
	}
	return result
}

func (c Coordinate) Move(dir string) Coordinate {
	switch dir {
	case "left":
		return c.Left()
	case "right":
		return c.Right()
	case "up":
		return c.Up()
	default:
		return c.Down()
	}
}

func (c Coordinate) Clone() Coordinate {
	result := Coordinate{
		X: c.X,
		Y: c.Y,
	}
	return result
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

type IndexResponse struct {
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
}

type Game struct {
	ID string `json:"id"`
}

type Board struct {
	Height int          `json:"height"`
	Width  int          `json:"width"`
	Food   []Coordinate `json:"food"`
	Snakes []Snake      `json:"snakes"`
}

func (b Board) Clone() Board {
	result := Board{
		Height: b.Height,
		Width:  b.Width,
		Food:   []Coordinate{},
		Snakes: []Snake{},
	}

	for _, coord := range b.Food {
		result.Food = append(result.Food, coord.Clone())
	}

	for _, snake := range b.Snakes {
		result.Snakes = append(result.Snakes, snake.Clone())
	}

	return result
}

type Snake struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Health int          `json:"health"`
	Body   []Coordinate `json:"body"`
	Shout  string       `json:"shout"`
}

func (s Snake) Clone() Snake {
	result := Snake{
		ID:     s.ID,
		Name:   s.Name,
		Health: s.Health,
		Body:   []Coordinate{},
		Shout:  s.Shout,
	}

	for _, coord := range s.Body {
		result.Body = append(result.Body, coord.Clone())
	}

	return result
}

// This is the request structure from the gameserver -
// We can pass the whole thing to any functions which are
// computing moves.
type MoveRequest struct {
	Game  `json:"game"`
	Turn  int `json:"turn"`
	Board `json:"board"`
	You   Snake `json:"you"`
}
