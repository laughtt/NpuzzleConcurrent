package pkg

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		fmt.Println("Bad file!")
		panic(e)
	}
}
func atoiAndCut(number *[]byte) int {
	var nbr int
	end := true

	for i := 0; i < len(*number); i++ {
		for (*number)[i] >= '0' && (*number)[i] <= '9' {
			n := int((*number)[i] - '0')
			nbr = nbr*10 + n
			i++
			end = false
		}
		if !end {
			*number = (*number)[i:]
			break
		}
	}
	return nbr
}

func ReadFile(path string) []int {
	dat, err := ioutil.ReadFile(path)
	check(err)
	//fmt.Println(string(dat))

	puzzleArray := make([]int, 0)

	for i := 0; i < len(dat); i++ {
		if dat[i] == '#' {
			endLine := i
			for len(dat) > endLine && dat[endLine] != '\n' {
				endLine++
			}
			dat = append(dat[0:i], dat[endLine:]...)
		}
	}

	repeatDcit := make(map[int]int, 0)
	puzzleSize := atoiAndCut(&dat)

	for i := 0; i < puzzleSize*puzzleSize; i++ {
		number := atoiAndCut(&dat)
		_, err := repeatDcit[number]
		if err {
			panic("Repetitive puzzle!")
		}
		repeatDcit[number]++
		if number >= puzzleSize*puzzleSize {
			panic("Bad puzzle!!!")
		}
		puzzleArray = append(puzzleArray, number)
	}
	return puzzleArray
}
