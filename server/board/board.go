package board

// 2d array by which the board is constructed
type base [48][128]Cell

// this struct has all information owned by the server
type Board struct {
	Base base
	Batallions []Batallion
	batallionMap map[[2]int]City
	Cities   	[]City
	cityMap  	map[[2]int]City
	Villages 	[]Village
	villageMap 	map[[2]int]Village

	//Player1 Player
	//Player2 Player
}
