package main

import (
	// "log"
	"cs302/final_project/Succession/server/board"
	"github.com/gdamore/tcell/v2"
	// "time"
)

func compare(b board.Batallion, c []board.City) int { // Helper function used to check if a batallion is on a city
	for i := 0; i < len(c); i++ {
		if b.X() == c[i].X() && b.Y() == c[i].Y() {
			return i
		}
	}
	return -1
}

func main() {
	var gameBoard board.Board = board.Generate_board() // Generating a board for gameplay usage
	screen := board.SetColors(&gameBoard)
	screen.Show()
	for { // main gameplay loop
		input := screen.PollEvent() // Takes the first event that happens to the screen
		switch input := input.(type) {
		case *tcell.EventKey: // We are only taking in keyboard inputs
			switch input.Rune() { // Classic switch statement logic with wasd movement
			case 'a':
				gameBoard.Batallions[0].Left()
				screen.Clear()
				result := compare(gameBoard.Batallions[0], gameBoard.Cities)
				if result != -1 {
					gameBoard.Cities = append(gameBoard.Cities[:result], gameBoard.Cities[result+1:]...)
				}
				screen = board.SetColors(&gameBoard)
				screen.Show()

			case 'w':
				gameBoard.Batallions[0].Up()
				screen.Clear()
				result := compare(gameBoard.Batallions[0], gameBoard.Cities)
				if result != -1 {
					gameBoard.Cities = append(gameBoard.Cities[:result], gameBoard.Cities[result+1:]...)
				}
				screen = board.SetColors(&gameBoard)
				screen.Show()

			case 's':

				gameBoard.Batallions[0].Down()
				screen.Clear()
				result := compare(gameBoard.Batallions[0], gameBoard.Cities)
				if result != -1 {
					gameBoard.Cities = append(gameBoard.Cities[:result], gameBoard.Cities[result+1:]...)
				}
				screen = board.SetColors(&gameBoard)
				screen.Show()

			case 'd':
				gameBoard.Batallions[0].Right()
				screen.Clear()

				result := compare(gameBoard.Batallions[0], gameBoard.Cities)
				if result != -1 {
					gameBoard.Cities = append(gameBoard.Cities[:result], gameBoard.Cities[result+1:]...)
				}
				screen = board.SetColors(&gameBoard)
				screen.Show()

			case 'q':
				return
			}
			if input.Key() == tcell.KeyCtrlC || input.Key() == tcell.KeyEscape {
				return
			}
			// time.Sleep(time.Millisecond)
		}
	}

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
