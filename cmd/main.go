package main

import (
	"fmt"

	"github.com/laughtt/NpuzzleConcurrent/pkg"
)

func executeOrder66() {

	//b := []int{5, 7, 1, 6, 0, 2, 4, 3, 8}

	//a := []int{1, 2, 3, 8, 0, 4, 7, 6, 5}

	c := []int{14,2,6,5,4,11,3,9,13 , 1 , 8 , 0 ,12 ,15 , 7 ,10}
	d := []int{1,2,3,4 , 12,13,14,5,11,0,15,6,10,9,8,7}
	mh := "ed"
	s := &pkg.Solver{}
	s.CreateSolver(c, d, mh)
	fmt.Println(s.Solve())
	// fmt.Println(h)
	// for h.dad != nil {
	// 	fmt.Printf("%d \n", h.mapa)
	// 	h = *h.dad
	// }
	// fmt.Printf("%d \n", h.mapa)
	// arrayPuzzles := createArrayPuzzle(&a)

}

func main() {
	executeOrder66()
}
