package entities

import (
	"math/rand"
)

// struct for Commander type
type Commander struct {
	name       string
	age        int
	attack     int
	defence    int
	open_field int
}

func (c *Commander) Name() string {
	return c.name
}
func (c *Commander) Age() int {
	return c.age
}
func (c *Commander) Attack() int {
	return c.attack
}
func (c *Commander) Defence() int {
	return c.defence
}
func (c *Commander) Open_fiend() int {
	return c.open_field
}

func Generate_commander() Commander {
	var c Commander
	c.name = ""
	c.age = 25
	c.attack = rand.Intn(10)
	c.defence = rand.Intn(10)
	c.open_field = rand.Intn(10)
	return c
}
