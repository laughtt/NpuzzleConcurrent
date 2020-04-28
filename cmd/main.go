package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/laughtt/NpuzzleConcurrent/pkg"
)

func printMapaSquare(puzzle []int) {
	side := math.Sqrt(float64(len(puzzle)))

	for i := 0; i < int(side); i++ {
		for ii := i * int(side); ii < (i+1)*int(side); ii++ {
			fmt.Printf("%d ", puzzle[ii])
		}
		fmt.Printf("\n")
	}
}

func executeOrder66(start []int, end []int, heuristic string) {

	//c := []int{14,2,6,5,4,11,3,9,13 , 1 , 8 , 0 ,12 ,15 , 7 ,10}
	//d := []int{1,2,3,4 , 12,13,14,5,11,0,15,6,10,9,8,7}

	s := &pkg.Solver{}
	s.CreateSolver(start, end, heuristic)

	if s.Solvable(start, end) {
		fmt.Println("This puzzle is not Solvable")
		return
	}
	solved := s.Solve()

	for solved != nil {
		printMapaSquare(solved.Mapa)
		fmt.Printf("\n")
		solved = solved.Dad
	}

	fmt.Println("OPEN:", s.Open)
	fmt.Println("MAX:", s.Close+s.Open)
	fmt.Println("CLOSE:", -s.Close)

}

func main() {
	final := flag.String("final", "3x3end", "final puzzle file name")
	start := flag.String("start", "3x3", "start puzzle file name")
	heuristic := flag.String("heuristic", "ed", "Heuristic algoritm")
	flag.Parse()
	pathFinal := "puzzles/" + *final
	pathStart := "puzzles/" + *start

	startPuzzle := pkg.ReadFile(pathStart)
	endPuzzle := pkg.ReadFile(pathFinal)
	printMapaSquare(startPuzzle)
	fmt.Println("\n")
	printMapaSquare(endPuzzle)
	fmt.Println("\n")
	executeOrder66(startPuzzle, endPuzzle, *heuristic)
}
