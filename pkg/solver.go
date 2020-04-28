package pkg

import (
	"fmt"
	"math"
	"sync"

	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

type cord struct {
	X int
	Y int
}

var moves = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type Q struct{}

type heuristic func(p1 *Puzzle, s *Solver) float32

type Solver struct {
	Algh        heuristic
	Dictionary  sync.Map
	Pq          *prque.Prque
	Start       *Puzzle
	FinalPuzzle *Puzzle
	Side        int
	Final       []cord
	Open        int
	Close       int
	Max         int
	Channel     chan *Puzzle
}

func createPuzzle(p []int, depth float32, score float32, dad *Puzzle, side int) *Puzzle {

	var X int
	var Y int

	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			X = i % side
			Y = i / side
			break
		}
	}

	return &Puzzle{
		Mapa:  p,
		Depth: depth,
		Score: score,
		Dad:   dad,
		Side:  side,
		X:     X,
		Y:     Y,
	}
}

func finalCoord(s []int) []cord {

	side := int(math.Sqrt(float64(len(s))))
	array := make([]cord, len(s))

	for i := 0; i < len(s); i++ {
		array[s[i]] = cord{
			X: i % side,
			Y: i / side,
		}
	}
	return array
}

func (s *Solver) CreateSolver(start []int, end []int, str string) {
	s.Algh = Heuristic(str)
	s.Dictionary = sync.Map{}
	s.Side = int(math.Sqrt(float64(len(start))))
	s.Start = createPuzzle(start, 0, 1, nil, s.Side)
	s.Final = finalCoord(end)
	s.Channel = make(chan *Puzzle)
	s.Pq = prque.New()
	s.Pq.Push(s.Start, 1)
	s.FinalPuzzle = createPuzzle(end, 0, 1, nil, s.Side)
	//fmt.Printf("%v", s.Algh(s.Start ,s))
}

func (s *Solver) Solvable(start []int, end []int) bool {
	invCount := 0
	for ii := 0; ii < s.Side*s.Side; ii += 1 {
		for jj := ii + 1; jj < s.Side*s.Side; jj += 1 {
			if start[ii] != 0 && start[jj] != 0 && start[ii] > start[jj] {
				invCount += 1
			}
		}
	}
	if s.Side%2 == 1 {
		return invCount%2 == 0
	} else if (((s.Side*s.Side)-1)/s.Side)%2 == 0 {
		return invCount%2 == 1
	} else {
		return invCount%2 == 0
	}
}

func copyPuzzle(p *Puzzle, X int, Y int) *Puzzle {
	mapa := make([]int, len(p.Mapa))

	copy(mapa, p.Mapa)

	newX, newY := (p.X + X), (p.Y + Y)
	mapa[p.X+p.Y*p.Side], mapa[newX+newY*p.Side] = mapa[newX+newY*p.Side], mapa[p.X+p.Y*p.Side]

	newPuzzle := Puzzle{
		Mapa:  mapa,
		Depth: p.Depth + 1,
		Dad:   p,
		Side:  p.Side,
		X:     newX,
		Y:     newY,
	}

	return &newPuzzle
}

func (s *Solver) checkDictionary(p *Puzzle) *Puzzle {
	str := fmt.Sprint(p.Mapa)

	if v, _ := s.Dictionary.Load(str); v == nil {
		s.Dictionary.Store(str, 1)
		return p
	}

	return nil
}

func applyHeuristic(p *Puzzle, s *Solver, X int, Y int) {

	if p.X+X >= s.Side || p.X+X < 0 || p.Y+Y >= s.Side || p.Y+Y < 0 {
		s.Channel <- nil
		return
	}

	newP := copyPuzzle(p, X, Y)

	if s.checkDictionary(newP) == nil {
		s.Channel <- nil
		return
	}

	newP.Score = s.Algh(newP, s)
	s.Channel <- newP
}

func (s *Solver) addPuzzles(p *Puzzle) {

	var a *Puzzle

	for i := range moves {
		go applyHeuristic(p, s, moves[i][0], moves[i][1])
	}

	for i := 0; i < 4; i++ {
		a = <-s.Channel
		if a != nil {
			s.Open++
			s.Max++

			s.Pq.Push(a, -float32(a.Score))

		}
	}
}

func (s *Solver) Solve() *Puzzle {

	ll := 0
	for !s.Pq.Empty() {
		p := s.Pq.PopItem()
		s.Close--
		s.Max--
		ll++
		if ll%100000 == 0 {
			fmt.Println(p)
		}
		puzzle := p.(*Puzzle)
		if puzzle.Score == 0 {
			return puzzle
		}

		s.addPuzzles(puzzle)
	}
	return nil
}
