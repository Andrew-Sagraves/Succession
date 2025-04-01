package render

import (
	"log"
	"fmt"
	"cs302/final_project/Succession/server/board"
	"github.com/gdamore/tcell/v2"
)


var colorBoard [48][128]tcell.Color

// func colorToANSI(color tcell.Color) string {
// 	r, g, b := color.RGB()
// 	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm]", r, g, b) // This provides a background color
// }

// Steps:
// Step 1: Take in the game board, and edit all of the game board indices with a color
//		- Go basic first, just make colorBoard land and water
//		- Assign specific colors to specific tiles
//		- Make biomes different shades
// Step 2: Try to layer a mouseinput board over the color board


func (b* board.Board.Base) setColors() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("failed to initialize screen: %v", err)
	}
	defer screen.Fini()
	// First thing we want to do is make specific colors for specific things
	// Loop through the original board
	for i := 0; i < 48; i++ {
		for j := 0; j < 128; j++ {
			if b[i][j].Foundation == board.Water { // Water is blue
				color := tcell.NewRGBColor(0, 0, 255)
				screen.SetContent(i, j, ' ', nil, tcell.StyleDefault.Background(color))
			} else { // Land is green
				color := tcell.NewRGBColor(0, 255, 0)
				screen.SetContent(i, j, ' ', nil, tcell.StyleDefault.Background(color))
			}
			// Need to make a priority list for the different types of cells
		}
	}
	screen.Show()
}

