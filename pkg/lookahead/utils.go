package lookahead

import "github.com/SamWheating/battlesnake2020/pkg/structs"

// check if a coordinate exists in a list
func CoordInList(c structs.Coordinate, list []structs.Coordinate) bool {
	for _, coord := range list {
		if coord == c {
			return true
		}
	}
	return false
}
