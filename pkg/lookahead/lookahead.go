package lookahead

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/SamWheating/battlesnake2020/pkg/heuristics"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

func Lookahead(state structs.MoveRequest, depth int, count int) string {
	directions := []string{"left", "right", "up", "down"}
	scores := make(map[string][]int)
	for _, direction := range directions {
		scores[direction] = []int{}
	}
	moves := SampleRandomSnakeMoves(state.Board, depth, count)
	results := make(map[string](chan int))
	var wg sync.WaitGroup

	results["left"] = make(chan int, count)
	results["right"] = make(chan int, count)
	results["up"] = make(chan int, count)
	results["down"] = make(chan int, count)

	for _, move := range moves {
		//score := scoreScenario(move, state, depth, channel) // adds (score, direction)
		wg.Add(1)
		go scoreScenario(move, state, depth, &wg, results)
	}
	wg.Wait()
	max := -10000.0
	choice := directions[rand.Int()%4] // a random default direction

	close(results["left"])
	close(results["right"])
	close(results["up"])
	close(results["down"])

	var log_line strings.Builder
	fmt.Fprintf(&log_line, "%d -%s\n", state.Turn, state.You.Name)

	for _, direction := range directions {
		count := len(results[direction])
		total := 0
		for i := range results[direction] {
			total += i
		}
		dirScore := float64(total) / float64(count)
		fmt.Fprintf(&log_line, "%s: %f\n", direction, dirScore)
		if dirScore > max {
			choice = direction
			max = dirScore
		}
	}
	fmt.Fprintf(&log_line, "chose: %s\n", choice)
	fmt.Println(log_line.String())
	return choice
}

// SampleRandomSnakeMoves generates <count> possible combinations of <depth> moves for each snake
// The moves aren't great but at the very least won't doubel back on itself or hit a wall
func SampleRandomSnakeMoves(board structs.Board, depth int, count int) []map[string][]string {
	rand.Seed(time.Now().Unix())
	snakemoves := make([]map[string][]string, count)
	directions := []string{"up", "down", "left", "right"}
	for i := 0; i < count; i++ {
		scenario := map[string][]string{}
		for _, snake := range board.Snakes {

			moves := make([]string, depth)
			head := structs.Coordinate{X: snake.Body[0].X, Y: snake.Body[0].Y}
			for j := 0; j < depth; j++ {

				bad_directions := make(map[string]bool) // just a set to keep track of unsafe moves

				// don't hit a wall
				if head.X == 0 {
					bad_directions["left"] = true
				}
				if head.X == board.Width-1 {
					bad_directions["right"] = true
				}
				if head.Up().Y < 0 || head.Up().Y >= board.Height {
					bad_directions["up"] = true
				}
				if head.Down().Y < 0 || head.Down().Y >= board.Height {
					bad_directions["down"] = true
				}

				// don't double back on ourselves
				if j > 1 {
					bad_directions[Opposite(moves[j-1])] = true
				}

				good_directions := []string{}
				for _, direction := range directions {
					if !bad_directions[direction] {
						good_directions = append(good_directions, direction)
					}
				}

				// only make a safe move
				direction := good_directions[rand.Intn(len(good_directions))]
				head = head.Move(direction)
				moves[j] = direction
			}
			scenario[snake.ID] = moves
		}
		snakemoves[i] = scenario
	}
	return snakemoves
}

func Opposite(move string) string {
	switch move {
	case "left":
		return "right"
	case "right":
		return "left"
	case "up":
		return "down"
	case "down":
		return "up"
	}
	return "error or something"
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
	newBoard := board.Clone()

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
	if snake.Health <= 0 { // Todo: this is a cheap hack for avoiding complete starvation
		return true
	}
	return false
}

func scoreScenario(moves map[string][]string, state structs.MoveRequest, depth int, wg *sync.WaitGroup, results map[string](chan int)) {
	//func scoreScenario(moves map[string][]string, state structs.MoveRequest, depth int, results map[string](chan int)) {
	// snake1: [left, right, down]
	newBoard := state.Board.Clone()
	direction := moves[state.You.ID][0]
	defer wg.Done()
	for i := 0; i < depth; i++ { // [left, right, down]
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
		// update snakes on the board to exclude dead snakes
		newBoard.Snakes = snakes

		// if we're dead by now, return score weighted by i
		alive := false
		for _, snake := range snakes {
			if snake.ID == state.You.ID {
				alive = true
				break
			}
		}
		if !alive {
			if i == 0 {
				results[direction] <- -100
				return
			}
			results[direction] <- -1 * (depth - i)
			return
		}
	}
	// if we made it all n turns, return the heuristic
	results[direction] <- heuristics.HeadRoom(newBoard, state.You.ID)
}
