package board

// 2d array by which the board is constructed
type base [48][128]Cell

// this struct has all information owned by the server
type Board struct {
	Base base

	Cities   []City
	Villages []Village
	Forts    []Fort

	Player1 Player
	Player2 Player
}
