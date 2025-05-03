package main

import (
	"fmt"
	// "log"
	"cs302/final_project/Succession/server/board"
	"os"
	"os/exec"
	// "github.com/gdamore/tcell/v2"
	// "time"
)

func main() {
	var gameBoard board.Board = board.Generate_board() // Generating a board for gameplay usage
	newScreen := board.SetColors(&gameBoard)
	// newScreen.Show()
	cmd := exec.Command("x-terminal-emulator", "-e", "tail -f /tmp/output.log")
	cmd.Start()
	file, err := os.Create("/tmp/output.log")
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer file.Close()
	for {
		newScreen.Show()
		// newScreen.EnableMouse()
		fmt.Println("Entered the loop")
		mouse := newScreen.PollEvent()
		switch mouse := mouse.(type) {
		case *tcell.EventMouse:
			x, y := mouse.Position()
			buttons := mouse.Buttons()
			_, err := fmt.Fprintf(file, "Mouse at (%d, %d) with buttons: %v\n", x, y, buttons)
			if err != nil {
				fmt.Println("Error writing to file:", err)
			}
			file.Sync()
		case *tcell.EventKey:
			if mouse.Key() == tcell.KeyEscape {
				return
			} else {
				fmt.Fprintf(file, "Key clicked!")
				file.Sync()
			}
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
