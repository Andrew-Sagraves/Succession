package board

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

func SetColors(b *Board) tcell.Screen {
	b.cityMap = make(map[[2]int]City)
	b.villageMap = make(map[[2]int]Village)
	b.batallionMap = make(map[[2]int]Batallion)
	for _, city := range b.Cities { // Sending all of the cities to a map for later use
		cityLocation := [2]int{city.X(), city.Y()}
		b.cityMap[cityLocation] = city
	}
	for _, village := range b.Villages { // Sending all of the cities to a map for later use
		villageLocation := [2]int{village.X(), village.Y()}
		b.villageMap[villageLocation] = village
	}
	for _, batallion := range b.Batallions {
		batallionLocation := [2]int {batallion.X(), batallion.Y()}
		b.batallionMap[batallionLocation] = batallion
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
	for i := 0; i < 48; i++ {
		for j := 0; j < 128; j++ {
			currentLocation := [2]int{i, j}
			if b.Base[i][j].Foundation == Water { // Water is blue
				color := tcell.NewRGBColor(0, 0, 255)
				screen.SetContent(j, i, ' ', nil, tcell.StyleDefault.Background(color))
			} else {
				if _, city := b.cityMap[currentLocation]; city { // Checks if a city exists in this specific vertex
					color := tcell.NewRGBColor(0, 0, 0)
					screen.SetContent(j, i, 'C', nil, tcell.StyleDefault.Background(color))
				} else if _, village := b.villageMap[currentLocation]; village { // Checks if a village exists in this specific vertex
					color := tcell.NewRGBColor(190, 100, 0)
					screen.SetContent(j, i, 'V', nil, tcell.StyleDefault.Background(color))
				} else if _, batallion := b.batallionMap[currentLocation]; batallion { // Checks if a batallion exists in this specific vertex
					color := tcell.NewRGBColor(0, 0, 128)
					screen.SetContent(j, i, 'B', nil, tcell.StyleDefault.Background(color))
				}else if b.Base[i][j].Biome == Agriculture { // This section is to make the map look a bit more colorful
					color := tcell.NewRGBColor(144, 238, 144)
					screen.SetContent(j, i, '*', nil, tcell.StyleDefault.Background(color))
				} else if b.Base[i][j].Biome == Field {
					color := tcell.NewRGBColor(83, 128, 89)
					screen.SetContent(j, i, ':', nil, tcell.StyleDefault.Background(color))
				} else if b.Base[i][j].Biome == Woods {
					color := tcell.NewRGBColor(63, 90, 54)
					screen.SetContent(j, i, '^', nil, tcell.StyleDefault.Background(color))
				}
			}
			// Need to make a priority list for the different types of cells
		}
	}
	return screen
}

// func UpdateColors(b *Board, screen tcell.Screen, previousX int, previousY int, currentX int, currentY int) tcell.Screen { A function meant to lower time taken to take input and display the change
// 	if b.Base[previousX][previousY].Foundation == Water { // Water is blue
// 		color := tcell.NewRGBColor(0, 0, 255)
// 		screen.SetContent(previousY, previousX, ' ', nil, tcell.StyleDefault.Background(color))
// 	} else {
// 		previousLocation := [2]int {previousX, previousY}
// 		if _, city := b.cityMap[previousLocation]; city { // Checks if a city exists in this specific vertex
// 			color := tcell.NewRGBColor(0, 0, 128) // Changes colors of the city to show the batallion has taken it
// 			screen.SetContent(previousY, previousX, 'C', nil, tcell.StyleDefault.Background(color))
// 		} else if _, village := b.villageMap[previousLocation]; village { // Checks if a village exists in this specific vertex
// 			color := tcell.NewRGBColor(0, 0, 128) // Changes colors of the city to show the batallion has taken it
// 			screen.SetContent(previousY, previousX, 'V', nil, tcell.StyleDefault.Background(color))
// 		}else if b.Base[previousX][previousY].Biome == Agriculture { // This section is to make the map look a bit more colorful
// 			color := tcell.NewRGBColor(144, 238, 144)
// 			screen.SetContent(previousY, previousX, '*', nil, tcell.StyleDefault.Background(color))
// 		} else if b.Base[previousX][previousY].Biome == Field {
// 			color := tcell.NewRGBColor(83, 128, 89)
// 			screen.SetContent(previousY, previousX, ':', nil, tcell.StyleDefault.Background(color))
// 		} else if b.Base[previousY][previousX].Biome == Woods {
// 			color := tcell.NewRGBColor(63, 90, 54)
// 			screen.SetContent(previousY, previousX, '^', nil, tcell.StyleDefault.Background(color))
// 		}
// 	}
// 	color := tcell.NewRGBColor(0,0,128) 
// 	screen.SetContent(currentY, currentX, 'B', nil, tcell.StyleDefault.Background(color)) This is to make the new cell the batallion cell
// 	return screen
// }
