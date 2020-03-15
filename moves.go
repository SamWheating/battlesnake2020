// Define move-computing functions in here to heep main.go clean

package main

func goLeft(state MoveRequest) string {
	return "left"
}

func goRight(state MoveRequest) string {
	return "right"
}

func followTail(MoveRequest) string {
	return "up"
}