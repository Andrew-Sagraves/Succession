package board

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
