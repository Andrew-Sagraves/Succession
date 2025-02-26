package entities

type Batallion struct {
	X           int
	Y           int
	Troop_Count int
	Allegiance  int
	Fitness     int
	Food        int
	Commander   Commander
}

func generate_batallion(x int, y int) Batallion {
	var c Batallion
	c.X = x
	c.Y = y
	c.Troop_Count = 0
	c.Allegiance = 0
	c.Food = 0
	c.Fitness = 0
	return c
}
