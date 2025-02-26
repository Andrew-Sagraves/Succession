package board

import (
	"math/rand"
)

// village struct
type Village struct {
	x          int
	y          int
	food       int
	population int
	allegiance int
}

// methods for village
func (v *Village) X() int {
	return v.x
}

func (c *Village) Y() int {
	return c.y
}
func (c *Village) Food() int {
	return c.food
}
func (c *Village) Population() int {
	return c.population
}
func (c *Village) Allegiance() int {
	return c.allegiance
}
func generate_village(x int, y int) Village {
	var c Village
	c.x = x
	c.y = y
	c.population = int(rand.Int63n(500) + 100)
	c.allegiance = 0
	return c
}

// fort struct
