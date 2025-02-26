package main

import (
	"succession/board"
	// "succession/render"
)

func main() {
	var b board.Board = board.Generate_board()

	board.Print_board_biome_test(b)
}

// for main file:
// have game loop that does certian things pass if time.NOW() >= time then + interval
