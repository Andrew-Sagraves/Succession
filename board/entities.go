package board

import (
	"math/rand"
)

type City struct {
	X          int
	Y          int
	Population int
	Allegiance int
}

func generate_city(x int, y int) City {
	var c City
	c.X = x
	c.Y = y
	c.Population = int(rand.Int63n(19000) + 1000)
	c.Allegiance = 0
	return c
}

type Village struct {
	X          int
	Y          int
	Population int
	Allegiance int
}

func generate_village(x int, y int) Village {
	var c Village
	c.X = x
	c.Y = y
	c.Population = int(rand.Int63n(500) + 100)
	c.Allegiance = 0
	return c
}

type Batallion struct {
	X           int
	Y           int
	Troop_Count int
	Allegiance  int
	Fitness     int
	Food        int
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
