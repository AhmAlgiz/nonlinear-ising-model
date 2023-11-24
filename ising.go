package main

import (
	"fmt"
	"math/rand"
)

type Ising struct {
	T   int
	J   int
	N   int
	Net [][]int
}

func (s *Ising) Fill() {
	s.Net = make([][]int, 0, s.N)
	for i := 0; i < s.N; i++ {
		s.Net[i] = make([]int, 0, s.N)
		for j := 0; j < s.N; j++ {
			s.Net[i][j] = rand.Intn(2)*2 - 1 // -1 or 1
		}
	}
	fmt.Print(s.Net)
}

func (s *Ising) Switch(i, j uint32) {

}

func main() {
	s := Ising{}
	s.Fill()
}
