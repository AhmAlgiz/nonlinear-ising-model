package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Ising struct {
	T   float64
	J   int
	N   int
	Kb  float64
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

func (s *Ising) P(e float64) float64 {
	return math.Exp((-1.0) * e / (s.Kb * s.T))
}

func (s *Ising) H(i, j int) int {
	if i == 0 {
		if j == 0 {
			return s.Net[i+1][j] + s.Net[i][j+1] + s.Net[s.N-1][j] + s.Net[i][s.N-1]
		}
		if j == s.N-1 {
			return s.Net[i][j-1] + s.Net[i+1][j] + s.Net[s.N-1][j] + s.Net[i][s.N-1]
		}
		return s.Net[i][j-1] + s.Net[i][j+1] + s.Net[s.N-1][j] + s.Net[i+1][j]
	}
	if i == s.N {
		if j == 0 {
			return s.Net[i-1][j] + s.Net[i][j+1] + s.Net[0][j] + s.Net[i][s.N-1]
		}
		if j == s.N-1 {
			return s.Net[i-1][j] + s.Net[i][j-1] + s.Net[0][j] + s.Net[i][0]
		}
		return s.Net[i][j-1] + s.Net[i][j+1] + s.Net[i-1][j] + s.Net[0][j]
	}
	return s.Net[i][j-1] + s.Net[i][j+1] + s.Net[i-1][j] + s.Net[i+1][j]
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

func (s *Ising) Switch() {

}

func main() {
	s := Ising{
		T:  2.0,
		J:  1,
		Kb: 1.0,
		N:  10,
	}
	s.Fill()
	fmt.Println(s.M())
	fmt.Println(s.H(0, 5))
}
