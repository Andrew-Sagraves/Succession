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
type Fort struct {
	x        int
	y        int
	food     int
	occupied int
}

// methods for fort
func (f *Fort) X() int {
	return f.x
}

func (f *Fort) Y() int {
	return f.y
}

func (f *Fort) Food() int {
	return f.food
}

func (f *Fort) Occupied() int {
	return f.occupied
}

func generate_fort(x int, y int, occupied int) Fort {
	var f Fort
	f.x = x
	f.y = y
	f.occupied = occupied
	return f
}
