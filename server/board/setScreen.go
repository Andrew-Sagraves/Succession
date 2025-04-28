package board

import  (
	"github.com/gdamore/tcell/v2"
	"log"
)

// var colorBoard [48][128]tcell.Color

func setColors(b *Board) {
	screen, err := tcell.NewScreen()
	if err != nil { // If screen is not initialized
		log.Fatal(err) 
	}
	defer screen.Fini() // Function used to clean up the screen
	err = screen.Init() // Initializes the display
	if err != nil {
		log.Fatal(err)
	}
	// First thing we want to do is make specific colors for specific things
	// Loop through the original board
	for i := 0; i < 48; i++ {
		for j := 0; j < 128; j++ {
			if b.Base[i][j].Foundation == Water { // Water is blue
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

// Steps:
// Step 1: Take in the game board, and edit all of the game board indices with a color
//		- Go basic first, just make colorBoard land and water
//		- Assign specific colors to specific tiles
//		- Make biomes different shades
// Step 2: Try to layer a mouseinput board over the color board



// Main Code for testing
// screen, err := tcell.NewScreen()
// 	if err != nil { // If screen is not initialized
// 		log.Fatal(err) 
// 	}
// 	defer screen.Fini() // Function used to clean up the screen
// 	err = screen.Init() // Initializes the display
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	running := true
// 	for running {
// 		// Draw Logic
// 		screen.Clear()

// 		screen.SetContent(0, 0, '@', nil, tcell.StyleDefault)

// 		screen.Show()



// 		ev := screen.PollEvent()

// 		switch ev := ev.(type) {
// 		case *tcell.EventKey:
// 			switch ev.Rune() {
// 			case 'q':
// 				running = false
// 			}
// 		}
// 	}

