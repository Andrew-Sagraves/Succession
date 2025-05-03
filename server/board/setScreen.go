package board

import  (
	"github.com/gdamore/tcell/v2"
	"log"
)

func SetColors (b *Board) tcell.Screen{
	cityMap := make(map[[2]int]City)
	villageMap := make(map[[2]int]Village)
	for _, city := range b.Cities { // Sending all of the cities to a map for later use
		cityLocation := [2]int{city.X(), city.Y()}
		cityMap[cityLocation] = city
	}
	for _, village := range b.Villages { // Sending all of the cities to a map for later use
		villageLocation := [2]int{village.X(), village.Y()}
		villageMap[villageLocation] = village
	}
	screen, err := tcell.NewScreen()
	if err != nil { // If screen is not initialized
		log.Fatal(err) 
	}
	// defer screen.Fini() // Function used to clean up the screen
	err = screen.Init() // Initializes the display
	if err != nil {
		log.Fatal(err)
	}
	// First thing we want to do is make specific colors for specific things
	// Loop through the original board
	for i := 0; i < 96; i +=2  {
		for j := 0; j < 128; j++ {
			currentLocation := [2]int{i / 2,j}
			if b.Base[i / 2][j].Foundation == Water { // Water is blue
				color := tcell.NewRGBColor(0, 0, 255)
				screen.SetContent(i, j, ' ', nil, tcell.StyleDefault.Background(color))
				screen.SetContent(i + 1, j, ' ', nil, tcell.StyleDefault.Background(color))
			} else { 
				if _, city := cityMap[currentLocation]; city { // Checks if a city exists in this specific vertex
					color := tcell.NewRGBColor(0, 0, 0)
					screen.SetContent(i, j, 'C', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, 'C', nil, tcell.StyleDefault.Background(color))
				}else if _, village := villageMap[currentLocation]; village{
					color := tcell.NewRGBColor(190,100,0)
					screen.SetContent(i, j, 'V', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, 'V', nil, tcell.StyleDefault.Background(color))
				}else if b.Base[i / 2][j].Biome == Agriculture {
					color := tcell.NewRGBColor(144, 238, 144)
					screen.SetContent(i, j, '*', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, '*', nil, tcell.StyleDefault.Background(color))

				} else if b.Base[i / 2][j].Biome == Field {
					color := tcell.NewRGBColor(83, 128, 89)
					screen.SetContent(i, j, ':', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, ':', nil, tcell.StyleDefault.Background(color))
				} else if b.Base[i / 2][j].Biome == Woods {
					color := tcell.NewRGBColor(63, 90, 54)
					screen.SetContent(i, j, '^', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, '^', nil, tcell.StyleDefault.Background(color))
				}
			}
			// Need to make a priority list for the different types of cells
		}
	}
	return screen
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

package board

import  (
	"github.com/gdamore/tcell/v2"
	"log"
)

func SetColors (b *Board) tcell.Screen{
	cityMap := make(map[[2]int]City)
	villageMap := make(map[[2]int]Village)
	for _, city := range b.Cities { // Sending all of the cities to a map for later use
		cityLocation := [2]int{city.X(), city.Y()}
		cityMap[cityLocation] = city
	}
	for _, village := range b.Villages { // Sending all of the cities to a map for later use
		villageLocation := [2]int{village.X(), village.Y()}
		villageMap[villageLocation] = village
	}
	screen, err := tcell.NewScreen()
	if err != nil { // If screen is not initialized
		log.Fatal(err) 
	}
	// defer screen.Fini() // Function used to clean up the screen
	err = screen.Init() // Initializes the display
	if err != nil {
		log.Fatal(err)
	}
	// First thing we want to do is make specific colors for specific things
	// Loop through the original board
	for i := 0; i < 96; i +=2  {
		for j := 0; j < 128; j++ {
			currentLocation := [2]int{i / 2,j}
			if b.Base[i / 2][j].Foundation == Water { // Water is blue
				color := tcell.NewRGBColor(0, 0, 255)
				screen.SetContent(i, j, ' ', nil, tcell.StyleDefault.Background(color))
				screen.SetContent(i + 1, j, ' ', nil, tcell.StyleDefault.Background(color))
			} else { 
				if _, city := cityMap[currentLocation]; city { // Checks if a city exists in this specific vertex
					color := tcell.NewRGBColor(0, 0, 0)
					screen.SetContent(i, j, 'C', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, 'C', nil, tcell.StyleDefault.Background(color))
				}else if _, village := villageMap[currentLocation]; village{
					color := tcell.NewRGBColor(190,100,0)
					screen.SetContent(i, j, 'V', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, 'V', nil, tcell.StyleDefault.Background(color))
				}else if b.Base[i / 2][j].Biome == Agriculture {
					color := tcell.NewRGBColor(144, 238, 144)
					screen.SetContent(i, j, '*', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, '*', nil, tcell.StyleDefault.Background(color))

				} else if b.Base[i / 2][j].Biome == Field {
					color := tcell.NewRGBColor(83, 128, 89)
					screen.SetContent(i, j, ':', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, ':', nil, tcell.StyleDefault.Background(color))
				} else if b.Base[i / 2][j].Biome == Woods {
					color := tcell.NewRGBColor(63, 90, 54)
					screen.SetContent(i, j, '^', nil, tcell.StyleDefault.Background(color))
					screen.SetContent(i + 1, j, '^', nil, tcell.StyleDefault.Background(color))
				}
			}
			// Need to make a priority list for the different types of cells
		}
	}
	return screen
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

