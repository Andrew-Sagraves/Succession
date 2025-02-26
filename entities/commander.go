package entities

type Commander struct {
	name       string
	age        int
	proficency int
}

func (c *Commander) Name() string {
	return c.name
}

func (c *Commander) Age() int {
	return c.age
}

func (c *Commander) Proficency() int {
	return c.proficency
}

func Generate_commander() Commander {
	var c Commander
	c.name = ""
	c.age = 25
	c.proficency = 0
	return c
}
