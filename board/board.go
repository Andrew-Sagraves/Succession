package board

type base [48][128]Cell

type Board struct {
	Base base

	Cities   []City
	Villages []Village
	Forts    []Fort
}
