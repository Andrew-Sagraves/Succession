package main

import (
	"succession/server/board"
	// "succession/render"
)

func main() {
	b := board.Generate_board()
	board.Print_board_biome_test(b)

	// this is all pseudocode, and feel free to put any testing code above, but this will outline the game
	// loop for the game...

	// first, we will have an enter() function, which will present the player's IP address, and have
	// the player type out the server ip address (perhaps their own) and the opponent IP address

	// second, there will be an initialize() function which is run by the server, generates the world,
	// and sends world info to each client

	// next, there will be a pick_city() function, which makes each player pick which city they want and
	// ends when both players have picked their initial city. this will have a game loop because information
	// must be communicated and rendered from the server to each client

	// last, we have the main game loop, which for each iteration will check for user input, then check for
	// gamestate from the server, and then will render the information (client). Alternatively, the server
	// will check for input from the client, and then respond with any needed data. this should be done
	// cocurrently
}
