package board

// base type
type Foundation int

const (
	Land Foundation = iota
	Water
)

// type of land/water
type Biome int

const (
	Field Biome = iota
	Woods
	Agriculture
	In_Water
)

// Cell combines Foundation and Biome into one type
type Cell struct {
	Foundation Foundation
	Biome      Biome
}
