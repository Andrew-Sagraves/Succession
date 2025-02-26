package entities

const messenger_speed = 10

type Messenger struct {
	x          int
	y          int
	allegiance int
	speed      int
}

func (m *Messenger) X() int {
	return m.x
}
func (m *Messenger) Y() int {
	return m.y
}
func (m *Messenger) Allegiance() int {
	return m.allegiance
}
func (m *Messenger) Speed() int {
	return m.speed
}

func generate_messenger(x int, y int, allegiance int) Messenger {
	var m Messenger
	m.x = x
	m.y = y
	m.allegiance = allegiance
	m.speed = messenger_speed
	return m
}
