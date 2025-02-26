package board

import (
	"math/rand"
)

// TODO: make member functions for each struct like change_allegiance, add_food, etc...
// make getters and setters
//

// city struct
type City struct {
	x          int
	y          int
	population int
	allegiance int
}

// methods for city
func (c *City) X() int {
	return c.x
}
func (c *City) Y() int {
	return c.y
}

func (c *City) Population() int {
	return c.population
}
func (c *City) Allegiance() int {
	return c.allegiance
}

func generate_city(x int, y int) City {
	var c City
	c.x = x
	c.y = y
	c.population = int(rand.Int63n(19000) + 1000)
	c.allegiance = 0
	return c
}
