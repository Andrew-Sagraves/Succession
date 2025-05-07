package board

import (
	"fmt"
	"math/rand"
)

var used [48][128]bool

// generates land/water for board struct
func recursive_generate_foundation(b *base, x int, y int, f Foundation) {
	used[x][y] = true

	// current ideal: .998, .997
	var land_percent = 0.997
	var water_percent = 0.997
	var probability = rand.Float64()

	// logic to determine what the cell is
	if f == Land {
		if land_percent > probability {
			b[x][y].Foundation = Land
		} else {
			b[x][y].Foundation = Water
		}
	} else {
		if water_percent > probability {
			b[x][y].Foundation = Water
		} else {
			b[x][y].Foundation = Land
		}
	}

	// logic to determine what cells to recursively call
	if rand.Int63n(2) == 1 {
		if (x-1 >= 0) && (!used[x-1][y]) {
			recursive_generate_foundation(b, x-1, y, b[x][y].Foundation)
		}

		if (y-1 >= 0) && (!used[x][y-1]) {
			recursive_generate_foundation(b, x, y-1, b[x][y].Foundation)
		}
		if (x+1 < 48) && (!used[x+1][y]) {
			recursive_generate_foundation(b, x+1, y, b[x][y].Foundation)
		}
		if (y+1 < 128) && (!used[x][y+1]) {
			recursive_generate_foundation(b, x, y+1, b[x][y].Foundation)
		}
	} else {
		if (y+1 < 128) && (!used[x][y+1]) {
			recursive_generate_foundation(b, x, y+1, b[x][y].Foundation)
		}
		if (x+1 < 48) && (!used[x+1][y]) {
			recursive_generate_foundation(b, x+1, y, b[x][y].Foundation)
		}
		if (y-1 >= 0) && (!used[x][y-1]) {
			recursive_generate_foundation(b, x, y-1, b[x][y].Foundation)
		}
		if (x-1 >= 0) && (!used[x-1][y]) {
			recursive_generate_foundation(b, x-1, y, b[x][y].Foundation)
		}

	}

}

// generates the biome. If the foundation is water, do not implement biome
func recursive_generate_biome(b *base, x int, y int, f Biome) {
	used[x][y] = true
	var ag_percent = 0.9
	var woods_percent = 0.8
	var field_percent = 0.8
	var probability = rand.Float64()

	if b[x][y].Foundation == Water {
		b[x][y].Biome = In_Water
		goto jump
	}

	// logic to determine what the cell is
	if f == Field {
		if field_percent > probability {
			b[x][y].Biome = Field
		} else {
			b[x][y].Biome = Woods
		}
	} else if f == Woods {
		if woods_percent > probability {
			b[x][y].Biome = Woods
		} else {
			b[x][y].Biome = Agriculture
		}
	} else {
		if ag_percent > probability {
			b[x][y].Biome = Agriculture
		} else {
			b[x][y].Biome = Field
		}

	}

jump:
	// logic to determine what cells to recursively call
	if rand.Int63n(2) == 1 {
		if (x+1 < 48) && (!used[x+1][y]) {
			recursive_generate_biome(b, x+1, y, b[x][y].Biome)
		}
		if (x-1 >= 0) && (!used[x-1][y]) {
			recursive_generate_biome(b, x-1, y, b[x][y].Biome)
		}
		if (y+1 < 128) && (!used[x][y+1]) {
			recursive_generate_biome(b, x, y+1, b[x][y].Biome)
		}
		if (y-1 >= 0) && (!used[x][y-1]) {
			recursive_generate_biome(b, x, y-1, b[x][y].Biome)
		}
	} else {
		if (y+1 < 128) && (!used[x][y+1]) {
			recursive_generate_biome(b, x, y+1, b[x][y].Biome)
		}
		if (y-1 >= 0) && (!used[x][y-1]) {
			recursive_generate_biome(b, x, y-1, b[x][y].Biome)
		}
		if (x+1 < 48) && (!used[x+1][y]) {
			recursive_generate_biome(b, x+1, y, b[x][y].Biome)
		}
		if (x-1 >= 0) && (!used[x-1][y]) {
			recursive_generate_biome(b, x-1, y, b[x][y].Biome)
		}

	}

}

// generates all cities and towns
func recursive_generate_populations(b *Board, x int, y int, n int) {
	used[x][y] = true
	var city_percent = 0.998
	var village_percent = 0.008
	var probability = rand.Float64()

	if b.Base[x][y].Foundation == Water {
		goto jump
	}

	// logic to determine what the cell is
	if n > 15 {
		if city_percent < probability {
			b.Cities = append(b.Cities, generate_city(x, y))
		}
		if village_percent > probability {
			b.Villages = append(b.Villages, generate_village(x, y))
		}
	}

jump:
	// logic to determine what cells to recursively call
	if rand.Int63n(2) == 1 {
		if (x+1 < 48) && (!used[x+1][y]) {
			recursive_generate_populations(b, x+1, y, n+1)
		}
		if (x-1 >= 0) && (!used[x-1][y]) {
			recursive_generate_populations(b, x-1, y, n+1)
		}
		if (y+1 < 128) && (!used[x][y+1]) {
			recursive_generate_populations(b, x, y+1, n+1)
		}
		if (y-1 >= 0) && (!used[x][y-1]) {
			recursive_generate_populations(b, x, y-1, n+1)
		}
	} else {
		if (y+1 < 128) && (!used[x][y+1]) {
			recursive_generate_populations(b, x, y+1, n+1)
		}
		if (y-1 >= 0) && (!used[x][y-1]) {
			recursive_generate_populations(b, x, y-1, n+1)
		}
		if (x+1 < 48) && (!used[x+1][y]) {
			recursive_generate_populations(b, x+1, y, n+1)
		}
		if (x-1 >= 0) && (!used[x-1][y]) {
			recursive_generate_populations(b, x-1, y, n+1)
		}

	}

}

// basic print function for debugging
func Print_board_foundation_test(b Board) {
	for i := 0; i < 48; i++ {
		for j := 0; j < 128; j++ {
			if b.Base[i][j].Foundation == Land {
				fmt.Print("L")
			} else {
				fmt.Print("W")
			}
		}
		fmt.Print("\n")
	}
}

// basic print for debugging
func Print_board_biome_test(b Board) {
	for i := 0; i < 48; i++ {
		for j := 0; j < 128; j++ {
			if b.Base[i][j].Biome == Agriculture {
				fmt.Print("A")
			} else if b.Base[i][j].Biome == Field {
				fmt.Print("F")
			} else if b.Base[i][j].Biome == Woods {
				fmt.Print("W")
			} else {
				fmt.Print("~")
			}
		}
		fmt.Print("\n")
	}
}

// base function that calls all recurive functions
func Generate_board() Board {

	var b Board

	var x int = rand.Intn(47)
	var y int = rand.Intn(127)

	b.Batallions = append(b.Batallions, generate_batallion(x, y, 1))


	// generate the foundation of the board
	b.Base[x][y].Foundation = Land
	recursive_generate_foundation(&b.Base, x, y, Land)

	// set used to false
	for i := range used {
		for j := range used[i] {
			used[i][j] = false
		}
	}
	// generate biomes
	b.Base[x][y].Biome = Field
	recursive_generate_biome(&b.Base, x, y, Field)

	// set used to false
	for i := range used {
		for j := range used[i] {
			used[i][j] = false
		}
	}

	// generate all cities/villages
	recursive_generate_populations(&b, x, y, 10)
	return b

}
