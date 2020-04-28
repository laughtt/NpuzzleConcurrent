package pkg

import "math"

func manhatanDistance(p1 *Puzzle, s *Solver) float32 {
	var dist float64

	pos := s.Final

	for i, v := range p1.Mapa {
		if v == 0 {
			continue
		}
		dist += math.Abs(float64(pos[v].X-i%s.Side) + math.Abs(float64(pos[v].Y-i/3)))
	}
	return float32(dist)
}

func titlesOutOfPlace(p1 *Puzzle, s *Solver) float32 {
	var dist float64

	pos := s.Final

	for i, v := range p1.Mapa {
		if v == 0 {
			continue
		}
		dist += math.Abs(float64(pos[v].X-i%3) + math.Abs(float64(pos[v].Y-i/3)))
	}
	return float32(dist)
}

func euclideanDistance(p1 *Puzzle, s *Solver) float32 {
	var dist float64

	pos := s.Final

	for i, v := range p1.Mapa {
		if v == 0 {
			continue
		}
		dist += math.Sqrt(math.Pow(float64(pos[v].X-i%3), 2) + math.Abs(math.Pow(float64(pos[v].Y-i/3), 2)))
	}
	return float32(dist)
}

//Heuristic a
func Heuristic(s string) func(p1 *Puzzle, s *Solver) float32 {
	switch s {
	case "mh":
		return manhatanDistance
	case "to":
		return titlesOutOfPlace
	case "ed":
		return euclideanDistance
	default:
		return manhatanDistance
	}
}
