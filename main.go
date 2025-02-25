package main

import (
	"fmt"
	"succession/board"
	// "succession/render"
)

func main() {
	var b board.Board = board.Generate_board()

	board.Print_board_biome_test(b)

	fmt.Println("villages:")

	for _, val := range b.Villages {
		fmt.Println(val)
	}
	fmt.Println(len(b.Villages))

	fmt.Println("cities:")

	for _, val := range b.Cities {
		fmt.Println(val)
	}
	fmt.Println(len(b.Cities))
}
