package entities

const batallion_speed = 5

type Batallion struct {
	x           int
	y           int
	troop_count int
	allegiance  int
	fitness     int
	food        int
	commander   Commander
	speed       int
}

func (b *Batallion) X() int {
	return b.x
}
func (b *Batallion) Troop_count() int {
	return b.troop_count
}
func (b *Batallion) Allegiance() int {
	return b.allegiance
}
func (b *Batallion) Fitness() int {
	return b.fitness
}
func (b *Batallion) Food() int {
	return b.food
}
func (b *Batallion) Commander() Commander {
	return b.commander
}
func (b *Batallion) Speed() int {
	return b.speed
}

func generate_batallion(x int, y int, allegiance int) Batallion {
	var c Batallion
	c.x = x
	c.y = y
	c.troop_count = 0
	c.allegiance = allegiance
	c.food = 0
	c.fitness = 0
	c.speed = batallion_speed
	return c
}
