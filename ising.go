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

func (s *Ising) M() int {
	m := 0
	for i := 0; i < s.N; i++ {
		for j := 0; j < s.N; j++ {
			m += s.Net[i][j]
		}
	}
	return m
}

func (s *Ising) Fill() {
	s.Net = make([][]int, 0, s.N)
	for i := 0; i < s.N; i++ {
		s.Net = append(s.Net, make([]int, s.N))
		for j := 0; j < s.N; j++ {
			s.Net[i][j] = rand.Intn(2)*2 - 1 // -1 or 1
		}
		fmt.Println(s.Net[i])
	}
}

func (s *Ising) Switch(i, j uint32) {

}

func main() {
	s := Ising{
		T: 2,
		J: 1,
		N: 10,
	}
	s.Fill()

}
