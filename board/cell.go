package board

// base type
type Foundation int

const (
	Land Foundation = iota
	Water
)

type Biome int

const (
	Field Biome = iota
	Woods
	Agriculture
	In_Water
)

type Cell struct {
	Foundation Foundation
	Biome      Biome
}
