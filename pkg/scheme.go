package pkg

type Puzzle struct {
	Mapa  []int
	Depth float32
	Score float32
	Dad   *Puzzle
	Side  int
	X     int
	Y     int
}
